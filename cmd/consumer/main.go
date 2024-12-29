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
)

func main() {
	godotenv.Load("./../../.env")
	topic := "websocket_gateway_to_notification_service"

	workersCount := 10
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
	})

	if err != nil {
		log.Fatalf("error on consumer creation: %s", err.Error())
	}

	defer consumer.Close()

	err = consumer.Subscribe(topic, nil)

	if err != nil {
		log.Fatalf("error on topic subscribing: %s", err.Error())
	}

	fmt.Println("the consumer is running...")

	// should ordering the messages before send to app client
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

	for w := 0; w < workersCount; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d shutting down...\n", workerId)
					return
				default:
					msg, err := consumer.ReadMessage(-1)
					if err != nil {
						fmt.Printf("\n error on reading message: %s", err.Error())
						continue
					}
					fmt.Printf("\n Received on Worker %d the message: %s ", workerId, string(msg.Value))
				}
			}
		}(w)
	}

	wg.Wait()
	fmt.Println("Shut down completed!")
	os.Exit(0)
}
