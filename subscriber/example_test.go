package subscriber_test

import (
	"fmt"
	"log"

	"github.com/dimmyjr/goku/message"
	"github.com/dimmyjr/goku/subscriber"
	"github.com/dimmyjr/goku/types"
)

func Example_minimal() {
	kafkaURL := "localhost:9092"
	topic := "topic-segmentio"
	groupID := "tests"

	consumer, err := subscriber.NewConsumer([]string{kafkaURL}, topic, groupID, types.Segmentio)

	if err != nil {
		log.Fatal("error to initialize producer")
	}

	consumer.Subscribe(read)
}

func read(msg *message.Message) error {
	fmt.Print(&msg.Value)

	return nil
}
