package consumer

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"strings"
)

type confluentConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  *kafka.Consumer
}

// NewConfluentConsumer Create new Confluent Consumer
func NewConfluentConsumer(kafkaURLs []string, topic, groupID string) (Consumer, error) {
	csm, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(kafkaURLs, ","),
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return confluentConsumer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		consumer:  csm,
	}, nil

}

func (con confluentConsumer) Subscribe(f func(message Message) error) {
	_ = con.consumer.SubscribeTopics([]string{con.topic, "^aRegex.*[Tt]opic"}, nil)
	go func() {
		for {
			msg, err := con.consumer.ReadMessage(-1)
			if err == nil {
				f(Message{
					Topic:     con.topic,
					Partition: msg.TopicPartition.Partition,
					Offset:    int64(msg.TopicPartition.Offset),
					Key:       msg.Key,
					Value:     msg.Value,
					//Headers:   msg.Headers,
					Time: msg.Timestamp,
				})
			} else {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()
}

func (con confluentConsumer) Close() {
	_ = con.consumer.Close()
}
