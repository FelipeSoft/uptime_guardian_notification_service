package main

// go env CGO_ENABLED
// 0
// go env -w CGO_ENABLED="1"
// go 1.21 for Confluent Apache Kafka v2

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FelipeSoft/uptime_guardian_notification_service/internal/infrastructure/kafka"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load("./../../.env")
	topic := "websocket_gateway_to_notification_service"

	producer := kafka.NewKafkaProducer()
	defer producer.Close()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal, shutting down...")
		cancel()
	}()

	fmt.Println("Kafka Producer is running...")

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			err := producer.Publish(&topic, "", "hello")
			if err != nil {
				fmt.Printf("error on producing message: %s", err.Error())
				continue
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer ctx.Done()
		<-ctx.Done()
		fmt.Println("Producer shutting down...")
		os.Exit(0)
	}()
	wg.Wait()

	fmt.Println("Shut down completed!")
}
