package producer

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type saramaProducer struct {
	kafkaURLs []string
	topic     string
	producer  sarama.SyncProducer
}

// NewSaramaProducer Create new Sarama Producer
func NewSaramaProducer(kafkaURLs []string, topic string) (Producer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Retry.Max = 5
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Return.Successes = true

	prd, err := sarama.NewSyncProducer(kafkaURLs, saramaCfg)

	if err != nil {
		return nil, err
	}

	return saramaProducer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		producer:  prd,
	}, nil

}

func (cfg saramaProducer) Publish(key, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: cfg.topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := cfg.producer.SendMessage(msg)

	if err != nil {
		return err
	}

	return nil
}

func (cfg saramaProducer) Close() {
	cfg.producer.Close()
}
