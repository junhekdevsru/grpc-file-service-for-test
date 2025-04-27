package main

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/junhekdevsru/grpc-file-service-for-test/proto"
)

func main() {
	// Подключаемся к серверу
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewFileServiceClient(conn)

	log.Println("Starting stress test on ListFiles...")

	const totalRequests = 150 // Генерируем 150 параллельных запросов

	var wg sync.WaitGroup
	wg.Add(totalRequests)

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		go func(id int) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := client.ListFiles(ctx, &proto.ListFilesRequest{})
			if err != nil {
				log.Printf("[Request #%d] Failed: %v", id, err)
			} else {
				log.Printf("[Request #%d] Success", id)
			}
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Completed %d requests in %s", totalRequests, elapsed)
}
