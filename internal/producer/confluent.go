package producer

import (
	"fmt"
	"strings"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type confluentProducer struct {
	kafkaURLs []string
	topic     string
	producer  *kafka.Producer
}

// NewConfluentProducer Create new Confluent Producer
func NewConfluentProducer(kafkaURLs []string, topic string) (Producer, error) {
	prc, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": strings.Join(kafkaURLs, ",")})
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range prc.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return confluentProducer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		producer:  prc,
	}, nil

}

func (prd confluentProducer) Publish(key, message string) error {
	msg := &kafka.Message{
		Key:            []byte(key),
		TopicPartition: kafka.TopicPartition{Topic: &prd.topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}
	err := prd.producer.Produce(msg, nil)
	prd.producer.Flush(100)

	if err != nil {
		return err
	}

	return nil
}

func (prd confluentProducer) Close() {
	prd.producer.Close()
}
