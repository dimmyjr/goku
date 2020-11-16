package consumer

import (
	"context"

	"github.com/dimmyjr/goku/message"
	"github.com/segmentio/kafka-go"
)

type SegmentioConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  *kafka.Reader
}

// NewSegmentioConsumer create new Segmentio Consumer.
func NewSegmentioConsumer(kafkaURLs []string, topic, groupID string) (*SegmentioConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURLs,
		GroupID:                groupID,
		Topic:                  topic,
		Partition:              0,
		Dialer:                 nil,
		QueueCapacity:          0,
		MinBytes:               0,
		MaxBytes:               0,
		MaxWait:                0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         0,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            0,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
	})

	return &SegmentioConsumer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		consumer:  reader,
	}, nil
}

func (segmentio SegmentioConsumer) Subscribe(f func(message *message.Message) error) {
	go segmentio.readMessages(f)
}

func (segmentio SegmentioConsumer) readMessages(f func(message *message.Message) error) {
	for {
		m, err := segmentio.consumer.ReadMessage(context.Background())
		if err != nil {
			break
		}

		_ = f(&message.Message{
			Topic:     m.Topic,
			Partition: int32(m.Partition),
			Offset:    m.Offset,
			Key:       m.Key,
			Value:     m.Value,
			Headers:   message.Headers(m.Headers),
			Time:      m.Time,
		})
	}
}

func (segmentio SegmentioConsumer) Close() {
	_ = segmentio.consumer.Close()
}
