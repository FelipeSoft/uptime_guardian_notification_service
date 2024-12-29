package kafka

import (
	// "context"
	// "net"
	// "os"

	// "github.com/confluentinc/confluent-kafka-go/kafka"
)

// type KafkaService struct {
// 	client 
// }

// func NewKafkaService(client *kafka.Client) *KafkaService {
// 	return &KafkaService{
// 		client: client,
// 	}
// }

// func (k *KafkaService) Publish(topic string) error {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	req := &kafka.ProduceRequest{
// 		Topic: topic,
// 		Addr: &net.IPAddr{
// 			IP: net.ParseIP(os.Getenv("KAFKA_BROKER")),
// 		},
// 		Partition: 1,
// 	}

// 	k.client.Produce(ctx, req)
// 	return nil
// }

// func (k *KafkaService) Consume(channel chan string) error {

// }

// func (k *KafkaService) Subscribe(ctx context.Context, req *kafka.JoinGroupRequest) error {
// 	_, err := k.client.JoinGroup(ctx, req)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
