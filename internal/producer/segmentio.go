package producer

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type SegmentioProducer struct {
	kafkaURLs []string
	topic     string
	producer  *kafka.Writer
}

// NewSegmentioProducer Create new Segmentio Producer.
func NewSegmentioProducer(kafkaURLs []string, topic string) (*SegmentioProducer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:           kafkaURLs,
		Topic:             topic,
		Dialer:            nil,
		Balancer:          &kafka.LeastBytes{},
		MaxAttempts:       0,
		QueueCapacity:     0,
		BatchSize:         0,
		BatchBytes:        0,
		BatchTimeout:      0,
		ReadTimeout:       0,
		WriteTimeout:      0,
		RebalanceInterval: 0,
		IdleConnTimeout:   0,
		RequiredAcks:      0,
		Async:             true,
		CompressionCodec:  nil,
		Logger:            nil,
		ErrorLogger:       nil,
	})

	return &SegmentioProducer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		producer:  writer,
	}, nil
}

func (segmentio SegmentioProducer) Publish(key, message string) error {
	msg := kafka.Message{
		Topic:     segmentio.topic,
		Partition: 0,
		Offset:    0,
		Key:       []byte(key),
		Value:     []byte(message),
		Headers:   nil,
		Time:      time.Time{},
	}
	err := segmentio.producer.WriteMessages(context.Background(), msg)
	//golint:gofumpt
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}

func (segmentio SegmentioProducer) Close() {
	if segmentio.producer != nil {
		_ = segmentio.producer.Close()
	}
}
