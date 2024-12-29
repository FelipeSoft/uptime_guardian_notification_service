package main

// go env CGO_ENABLED
// 0
// go env -w CGO_ENABLED="1"
// go 1.21 for Confluent Apache Kafka v2

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load("./../../.env")
	topic := "websocket_gateway_to_notification_service"

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
	})

	if err != nil {
		log.Fatalf("error on producer creation: %s", err.Error())
	}
	defer producer.Close()

	fmt.Println("the producer is running...")

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

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            nil,
				Value:          []byte(fmt.Sprintf("hello from message %d", i)),
			}, nil)
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
	}()
	wg.Wait()

	fmt.Println("Flushing pending messages...")
	unflushed := producer.Flush(15 * 1000)
	if unflushed > 0 {
		fmt.Printf("Warning: %d messages were not flushed\n", unflushed)
	} else {
		fmt.Println("All messages flushed successfully.")
	}

	fmt.Println("Shut down completed!")
}
