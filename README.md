[![PkgGoDev](https://pkg.go.dev/badge/github.com/dimmyjr/goku)](https://pkg.go.dev/github.com/dimmyjr/goku)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimmyjr/goku)](https://goreportcard.com/report/github.com/dimmyjr/goku)

![logo](https://user-images.githubusercontent.com/3587262/93725006-1ed7c980-fb82-11ea-9bcc-cfb33cd4bbbf.jpg)

Integrate golang with kafka in a simple and fast way, being able to choose the provider you want.

# Features
- Producer
- Consumer
- gRPC Server
- Metrics on Prometheus and Grafana

# Providers
- https://github.com/segmentio/kafka-go
- https://github.com/Shopify/sarama
- https://github.com/confluentinc/confluent-kafka-go

# Usage

### Publisher (Producer)
```
package publisher_test

import (
	"github.com/dimmyjr/goku/publisher"
	"github.com/dimmyjr/goku/types"
	"log"
)

func Example_minimal() {
	kafkaURL := "localhost:9092"
	topic := "topic-segmentio"
	key := "topicKey"
	message := " ...kafka message... "

	producer, err := publisher.New([]string{kafkaURL}, topic, types.Segmentio)

	if err != nil {
		log.Fatal("error to initialize producer")
	}

	err = producer.Publish(key, message)

	if err != nil {
		log.Fatal("error to send message")
	}
}
```

### Subscriber (Consumer)
```
package subscriber_test

import (
	"fmt"
	"github.com/dimmyjr/goku/subscriber"
	"github.com/dimmyjr/goku/types"
	"log"
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

func read(input subscriber.Message) error {
	fmt.Print(input.Value)
	return nil
}

```

# Contributing
Interested in contributing to GoKu? We'd love your help. GoKu is an open source project, built one contribution at a time by users like you. See the CONTRIBUTING file for instructions on how to contribute.

# References
- https://www.confluent.io/blog/exactly-once-semantics-are-possible-heres-how-apache-kafka-does-it/
- https://github.com/edenhill/librdkafka
