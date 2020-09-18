package consumer

import (
	"fmt"
	"strings"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type confluentConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  *kafka.Consumer
}

// NewConfluentConsumer Create new Confluent Consumer
func NewConfluentConsumer(kafkaURLs []string, topic, groupID string) (Consumer, error) {
	//prc, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": strings.Join(kafkaURLs, ",")})
	//if err != nil {
	//	return nil, err
	//}
	//
	//go func() {
	//	for e := range prc.Events() {
	//		switch ev := e.(type) {
	//		case *kafka.Message:
	//			if ev.TopicPartition.Error != nil {
	//				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	//			}
	//		}
	//	}
	//}()
	//
	//return confluentConsumer{
	//	kafkaURLs: kafkaURLs,
	//	topic:     topic,
	//	consumer:  prc,
	//}, nil
	return nil, nil
}

func (con segmentioConsumer) Subscribe(f func(message Message) error) {
	//msg := &kafka.Message{
	//	Key:            []byte(key),
	//	TopicPartition: kafka.TopicPartition{Topic: &prd.topic, Partition: kafka.PartitionAny},
	//	Value:          []byte(message),
	//}
	//err := prd.consumer.Produce(msg, nil)
	//prd.consumer.Flush(100)

}

func (prd confluentConsumer) Close() {
	_ = prd.consumer.Close()
}
