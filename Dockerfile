# 1: сборка приложения
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Статическая сборка под Linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# 2: чистый контейнер
FROM debian:bullseye-slim

# Устанавливаем сертификаты (если будем использовать TLS в будущем)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Копируем бинарник из билдера
COPY --from=builder /app/server .

# Копируем proto папку (если нужно для ошибок и совместимости)
COPY --from=builder /app/proto ./proto

# Открываем порт для gRPC сервера
EXPOSE 50051

# Запуск приложения
CMD ["./server"]
