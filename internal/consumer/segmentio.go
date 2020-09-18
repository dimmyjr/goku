package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type segmentioConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  *kafka.Reader
}

// NewSegmentioConsumer Create new Segmentio Consumer
func NewSegmentioConsumer(kafkaURLs []string, topic, groupID string) (Consumer, error) {
	rdr := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        kafkaURLs,
		GroupID:        groupID,
		Topic:          topic,
		Logger:         nil,
		ErrorLogger:    nil,
		IsolationLevel: 0,
		MaxAttempts:    0,
	})

	return segmentioConsumer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		consumer:  rdr}, nil

}

func (con segmentioConsumer) Subscribe(f func(message Message) error) {
	for {
		m, err := con.consumer.ReadMessage(context.Background())
		if err != nil {
			break
		}
		_ = f(Message{
			Topic:     m.Topic,
			Partition: m.Partition,
			Offset:    m.Offset,
			Key:       m.Key,
			Value:     m.Value,
			//Headers:   m.Headers,
			Time: m.Time,
		})
	}
}

func (con segmentioConsumer) Close() {
	_ = con.consumer.Close()
}
