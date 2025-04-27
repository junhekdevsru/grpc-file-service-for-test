package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/junhekdevsru/grpc-file-service-for-test/internal/service"
	"github.com/junhekdevsru/grpc-file-service-for-test/proto"
)

func main() {
	// Путь к хранению файлов
	storagePath := "./storage"

	// Проверяем существует ли папка, если нет - создаём
	if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
		log.Fatalf("failed to create storage directory: %v", err)
	}

	// Создаем TCP-листнер на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаём grpc-сервер
	grpcServer := grpc.NewServer()

	// Создаём экземпляр FileService
	fileService := service.NewFileService(storagePath)

	// Регистрируем сервис в grpc-сервере
	proto.RegisterFileServiceServer(grpcServer, fileService)

	// Регистрируем reflection для grpcurl и др. клиентов
	reflection.Register(grpcServer)

	log.Println("grpc server listening on :50051")

	// Запуск сервера
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
