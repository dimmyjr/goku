package producer

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

type segmentioProducer struct {
	kafkaURLs []string
	topic     string
	producer  *kafka.Writer
}

// NewSegmentioProducer Create new Segmentio Producer
func NewSegmentioProducer(kafkaURLs []string, topic string) (Producer, error) {
	prc := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaURLs,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return segmentioProducer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		producer:  prc}, nil

}

func (prd segmentioProducer) Publish(key, message string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	}
	err := prd.producer.WriteMessages(context.Background(), msg)

	if err != nil {
		return err
	}

	return nil
}

func (prd segmentioProducer) Close() {
	prd.producer.Close()
}
