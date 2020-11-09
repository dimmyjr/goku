package producer

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

type SaramaProducer struct {
	kafkaURLs []string
	topic     string
	producer  sarama.SyncProducer
}

// NewSaramaProducer Create new Sarama Producer.
func NewSaramaProducer(kafkaURLs []string, topic string) (*SaramaProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Retry.Max = 5
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Return.Successes = true

	prd, err := sarama.NewSyncProducer(kafkaURLs, saramaCfg)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &SaramaProducer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		producer:  prd,
	}, nil
}

func (saramaProducer SaramaProducer) Publish(key, message string) error {
	msg := &sarama.ProducerMessage{
		Topic:     saramaProducer.topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.StringEncoder(message),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}
	_, _, err := saramaProducer.producer.SendMessage(msg)
	//golint:gofumpt
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}

func (saramaProducer SaramaProducer) Close() {
	if saramaProducer.producer != nil {
		_ = saramaProducer.producer.Close()
	}
}
