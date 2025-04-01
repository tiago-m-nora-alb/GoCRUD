package kafkaConfig

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

var (
	producerInstance *Producer
	onceProducer     sync.Once
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	var err error
	onceProducer.Do(func() {
		p, e := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": os.Getenv("KAFKA_SERVERS"),
		})
		if e != nil {
			err = e
			return
		}

		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						log.Printf("Falha ao entregar mensagem: %v\n", ev.TopicPartition.Error)
					}
				}
			}
		}()

		producerInstance = &Producer{producer: p}
	})
	return producerInstance, err
}

func (p *Producer) SendMessage(topic string, key string, value interface{}) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          payload,
	}, nil)
}

func (p *Producer) Close() {
	p.producer.Flush(15 * 1000)
	p.producer.Close()
}
