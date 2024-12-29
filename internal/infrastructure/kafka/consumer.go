package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer() *KafkaConsumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
	})
	if err != nil {
		log.Fatalf("error on consumer creation %s", err.Error())
	}
	return &KafkaConsumer{
		consumer: consumer,
	}
}

func (c *KafkaConsumer) Subscribe(topic string) {
	err := c.consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("error on subscribing topic: %s", err.Error())
	}
}

func (c *KafkaConsumer) Consume() (*kafka.Message, error) {
	msg, err := c.consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *KafkaConsumer) Close() {
	fmt.Println("connection closed with consumer...")
	c.consumer.Close()
}