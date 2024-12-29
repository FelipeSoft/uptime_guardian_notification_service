package kafka

import (
	"log"
	"os"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() *KafkaProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
	})
	if err != nil {
		log.Fatalf("error on producer creation %s", err.Error())
	}
	return &KafkaProducer{
		producer: producer,
	}
}

func (p *KafkaProducer) Publish(topic *string, key, message string) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *KafkaProducer) Close() {
	p.producer.Close()
}
