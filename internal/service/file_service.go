package service

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/junhekdevsru/grpc-file-service-for-test/proto"
)

type FileService struct {
	proto.UnimplementedFileServiceServer // реализация grpc-сервера

	uploadLimiter chan struct{} // семафор на upload/download (10 слотов)
	listLimiter   chan struct{} // семафор на list (100 слотов)
	storagePath   string        // путь к папке хранения файлов
}

// Конструктор нового FileService
func NewFileService(storagePath string) *FileService {
	return &FileService{
		uploadLimiter: make(chan struct{}, 10),
		listLimiter:   make(chan struct{}, 100),
		storagePath:   storagePath,
	}
}

// 		=#	Заготовки методов	#=

// Загрузка файла (stream от клиента)
func (s *FileService) UploadFile(stream proto.FileService_UploadFileServer) error {
	// TODO: логика обработки загрузки
	// Захватываем слот для ограничения подключения
	s.uploadLimiter <- struct{}{}
	defer func() {
		<-s.uploadLimiter
	}()

	var fileName string
	var file *os.File

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break // Клиент закончил отправлять данные
		}
		if err != nil {
			return err // Ошибка в стриме
		}
		// Файл открываем только при первом чанке
		if file == nil {
			fileName = req.GetFilename()
			if fileName == "" {
				return status.Error(codes.InvalidArgument, "filename is required in the first chunk")
			}

			filePath := filepath.Join(s.storagePath, fileName)
			file, err = os.Create(filePath)
			if err != nil {
				return status.Errorf(codes.Internal, "file to create file: %v", err)
			}
			defer file.Close()
		}

		// Записываем контент чанка в файл
		if _, err := file.Write(req.GetContent()); err != nil {
			return status.Errorf(codes.Internal, "failed to write to file: %v", err)
		}
	}
	// Отправляем ответ об успешной загрузке
	return stream.SendAndClose(&proto.UploadFileResponse{
		Message: "file uploaded seccessfully",
	})
}

// Скачивание файла (stream к клиенту)
func (s *FileService) ListFiles(ctx context.Context, req *proto.ListFilesRequest) (*proto.ListFilesResponse, error) {
	//  Захватываем слот для ограничения подключений
	s.listLimiter <- struct{}{}
	defer func() {
		<-s.listLimiter
	}()

	files, err := os.ReadDir(s.storagePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read storage directory: %v", err)
	}

	var fileInfos []*proto.FileInfo

	for _, file := range files {
		if file.IsDir() {
			continue // пропускаем папки
		}

		filePath := filepath.Join(s.storagePath, file.Name())

		stat, err := os.Stat(filePath)
		if err != nil {
			continue // пропускаем файлы, к которым нет доступа
		}

		fileInfos = append(fileInfos, &proto.FileInfo{
			Name:      file.Name(),
			CreatedAt: stat.ModTime().Format(time.RFC3339), // на Unix-системах нет отдельного создания
			UpdatedAt: stat.ModTime().Format(time.RFC3339),
		})
	}

	return &proto.ListFilesResponse{
		Files: fileInfos,
	}, nil
}

// DownloadFile реализует скачивание файла через стрим
func (s *FileService) DownloadFile(req *proto.DownloadFileRequest, stream proto.FileService_DownloadFileServer) error {
	// Захватываем слот для ограничения подключений
	s.uploadLimiter <- struct{}{}
	defer func() {
		<-s.uploadLimiter
	}()

	filePath := filepath.Join(s.storagePath, req.GetFilename())

	file, err := os.Open(filePath)
	if err != nil {
		return status.Errorf(codes.NotFound, "file not found: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1MB буфер для отправки чанков

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return status.Errorf(codes.Internal, "failed to read file: %v", err)
		}
		if n == 0 {
			break // конец файла
		}

		res := &proto.DownloadFileResponse{
			Content: buffer[:n],
		}

		if err := stream.Send(res); err != nil {
			return status.Errorf(codes.Internal, "failed to send chunk: %v", err)
		}
	}

	return nil
}
