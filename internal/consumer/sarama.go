package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

type saramaConsumer struct {
	kafkaURLs []string
	topic     string
	consumer  sarama.ConsumerGroup
}

// ConsumerGroupHandler represents the sarama consumer group
type GroupHandler struct {
	f func(message Message) error
}

func (h GroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h GroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages(), here is supposed to be what you want to
// do with the message. In this example the message will be logged with the topic name, partition and message value.
func (h GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		_ = h.f(Message{
			Topic:     msg.Topic,
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Key:       msg.Key,
			Value:     msg.Value,
			//Headers:   nil,
			Time: msg.Timestamp,
		})
		session.MarkMessage(msg, "")
	}

	return nil
}

// NewSaramaConsumer Create new Sarama Consumer
func NewSaramaConsumer(kafkaURLs []string, topic, groupID string) (Consumer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	con, err := sarama.NewConsumerGroup(kafkaURLs, groupID, config)
	if err != nil {
		return nil, err
	}

	return saramaConsumer{
		kafkaURLs: kafkaURLs,
		topic:     topic,
		consumer:  con,
	}, nil
}

func (con saramaConsumer) Subscribe(f func(message Message) error) {
	go func() {
		for {
			_ = con.consumer.Consume(context.Background(), []string{con.topic}, GroupHandler{f})
		}
	}()
}

func (con saramaConsumer) Close() {
	_ = con.consumer.Close()
}
