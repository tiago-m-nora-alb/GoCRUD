package kafkaConfig

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func CreateTopics() {
	createTopic("products")
}

func createTopic(topic string) {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVERS"),
	})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	defer admin.Close()

	newTopic := topic
	topicSpec := kafka.TopicSpecification{
		Topic:             newTopic,
		NumPartitions:     1, 
		ReplicationFactor: 1,
	}

	timeout := 60 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	results, err := admin.CreateTopics(ctx, []kafka.TopicSpecification{topicSpec}, kafka.SetAdminOperationTimeout(timeout))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Topic %s creation failed: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s created successfully\n", result.Topic)
		}
	}
}
