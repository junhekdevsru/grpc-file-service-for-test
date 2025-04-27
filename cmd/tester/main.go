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
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewFileServiceClient(conn)

	log.Println("Starting stress test...")

	// Стресс-тест на ListFiles (ограничение 100 одновременных запросов)
	testListFiles(client, 150)
}

func testListFiles(client proto.FileServiceClient, totalRequests int) {
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := client.ListFiles(ctx, &proto.ListFilesRequest{})
			if err != nil {
				log.Printf("[Request %d] Error: %v", id, err)
			} else {
				log.Printf("[Request %d] Success", id)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	log.Printf("Completed %d requests in %s", totalRequests, elapsed)
}
