package consumer

import (
	"github.com/Shopify/sarama"
)

type saramaConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  sarama.Consumer
}

// NewSaramaConsumer Create new Sarama Consumer
func NewSaramaConsumer(kafkaURLs []string, topic, groupID string) (Consumer, error) {
	//sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	//
	//saramaCfg := sarama.NewConfig()
	//saramaCfg.Consumer.Retry.Max = 5
	//saramaCfg.Consumer.RequiredAcks = sarama.WaitForAll
	//saramaCfg.Consumer.Return.Successes = true
	//
	//prd, err := sarama.NewSyncConsumer(kafkaURLs, saramaCfg)
	//
	//if err != nil {
	//	return nil, err
	//}

	//return saramaConsumer{
	//	kafkaURLs: kafkaURLs,
	//	topic:     topic,
	//	consumer:  prd,
	//}, nil

	return nil, nil
}

func (con saramaConsumer) Subscribe(f func(message Message) error) {
	//msg := &sarama.ConsumerMessage{
	//	Topic: cfg.topic,
	//	Value: sarama.StringEncoder(message),
	//}
	//_, _, err := cfg.consumer.SendMessage(msg)
	//
	//if err != nil {
	//	return err
	//}

}

func (con saramaConsumer) Close() {
	//	con.consumer.Close()
}
