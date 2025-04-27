package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc-file-service/proto" // Путь пока локальный, потом поменяем на github.com/...
)

func main() {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Отключаем TLS для локальной разработки
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Создаём клиента FileService
	client := proto.NewFileServiceClient(conn)

	// Имя файла для загрузки/скачивания
	filename := "test_upload.txt"

	// Загружаем файл на сервер
	if err := uploadFile(client, filename); err != nil {
		log.Fatalf("upload failed: %v", err)
	}

	// Скачиваем файл с сервера
	if err := downloadFile(client, filename); err != nil {
		log.Fatalf("download failed: %v", err)
	}
}

// uploadFile загружает файл на сервер через UploadFile RPC
func uploadFile(client proto.FileServiceClient, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create upload stream: %w", err)
	}

	buffer := make([]byte, 1024*1024) // Буфер 1MB
	firstChunk := true

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			break
		}

		req := &proto.UploadFileRequest{
			Content: buffer[:n],
		}
		if firstChunk {
			req.Filename = filename
			firstChunk = false
		}

		if err := stream.Send(req); err != nil {
			return fmt.Errorf("failed to send chunk: %w", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive upload response: %w", err)
	}

	log.Println("[UPLOAD SUCCESS]:", res.Message)
	return nil
}

// downloadFile скачивает файл с сервера через DownloadFile RPC
func downloadFile(client proto.FileServiceClient, filename string) error {
	req := &proto.DownloadFileRequest{
		Filename: filename,
	}

	stream, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to start download: %w", err)
	}

	outFile, err := os.Create("downloaded_" + filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive chunk: %w", err)
		}

		if _, err := outFile.Write(res.GetContent()); err != nil {
			return fmt.Errorf("failed to write chunk to file: %w", err)
		}
	}

	log.Println("[DOWNLOAD SUCCESS]:", "File saved as downloaded_"+filename)
	return nil
}
