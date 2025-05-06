# gRPC File Service

## 📋 Описание задачи

- Сервис принимает бинарные файлы от клиента и сохраняет их на диск.
- Позволяет получить список всех загруженных файлов (имя, дата создания, дата обновления).
- Позволяет скачать загружённые файлы обратно.
- Ограничивает количество одновременных подключений:
  - На загрузку/скачивание файлов — до 10 конкурентных запросов.
  - На просмотр списка файлов — до 100 конкурентных запросов.

## ⚙️ Стек технологий

- Go 1.24.2
- gRPC
- Protocol Buffers
- Docker

##  Структура проекта
grpc-file-service/
├── cmd/
│   └── server/
│       └── main.go                # Точка входа: запуск сервера, инициализация зависимостей
├── internal/
│   ├── domain/
│   │   └── file/
│   │       ├── entity.go          # Структура File: имя, путь, дата, и т.д.
│   │       └── repository.go      # Интерфейс FileRepository
│   ├── usecase/
│   │   └── file/
│   │       └── usecase.go         # Логика: UploadFile, ListFiles, DownloadFile
│   ├── delivery/
│   │   └── grpc/
│   │       └── handler.go         # Реализация gRPC-сервиса: FileServiceServer
│   ├── infrastructure/
│   │   ├── limiter/
│   │   │   └── semaphore.go       # Семафор с учётом context
│   │   ├── storage/
│   │   │   └── local_fs.go        # Реализация репозитория на локальной ФС
│   │   └── logger/
│   │       └── logger.go          # Логгер с уровнями
├── proto/
│   └── file_service.proto         # gRPC-описание
├── configs/
│   └── config.yaml                # Конфиг сервера: порт, пути и т.п.
├── scripts/
│   └── test_stress.sh             # Стресс-тест на лимитеры
├── client/
│   └── main.go                    # gRPC-клиент для загрузки/скачивания
├── Dockerfile
├── go.mod
├── README.md


### Сборка и запуск через Docker

```bash
docker build -t grpc-file-service .
docker run -p 50051:50051 grpc-file-service

##  Тестирование

### Быстрая проверка сервиса

```bash
cd cmd/tester
go run main.go
