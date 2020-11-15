package consumer

import (
	"fmt"
	"strings"

	"github.com/dimmyjr/goku/message"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ConfluentConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  *kafka.Consumer
}

// NewConfluentConsumer create new Confluent Consumer.
func NewConfluentConsumer(kafkaURLs []string, topic, groupID string) (*ConfluentConsumer, error) {
	csm, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(kafkaURLs, ","),
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &ConfluentConsumer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		consumer:  csm,
	}, nil
}

func (confluent ConfluentConsumer) Subscribe(f func(message *message.Message) error) {
	_ = confluent.consumer.SubscribeTopics([]string{confluent.topic, "^aRegex.*[Tt]opic"}, nil)

	go confluent.readMessages(f)
}

func (confluent ConfluentConsumer) readMessages(f func(message *message.Message) error) {
	for {
		msg, err := confluent.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		if f != nil {
			err = f(&message.Message{
				Topic:     confluent.topic,
				Partition: msg.TopicPartition.Partition,
				Offset:    int64(msg.TopicPartition.Offset),
				Key:       msg.Key,
				Value:     msg.Value,
				Headers:   message.Headers(msg.Headers),
				Time:      msg.Timestamp,
			})
		}

		if err != nil {
			fmt.Printf("Consumer callback error: %v (%v)\n", err, msg)
		}
	}
}

func (confluent ConfluentConsumer) Close() {
	_ = confluent.consumer.Close()
}
