// Package subscriber implements Kafka Consumer for Go.
package subscriber

import (
	"errors"
	"fmt"
	"time"

	"github.com/dimmyjr/goku/internal/consumer"
	"github.com/dimmyjr/goku/message"
	"github.com/dimmyjr/goku/types"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// ErrUndefinedProvider : undefined provider error.
var ErrUndefinedProvider = errors.New("undefinedProvider")

type Subscriber struct {
	histogram *prometheus.HistogramVec
	provider  types.Provider
	consumer  consumer.Consumer
}

func initPrometheus() *prometheus.HistogramVec {
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem:   "Subscriber",
			ConstLabels: prometheus.Labels{},
			Namespace:   "kafka_consumer",
			Name:        "consumer",
			Help:        "Consumer execution in seconds",
			Buckets:     prometheus.DefBuckets,
		}, []string{"name"})

	if histogram != nil {
		_ = prometheus.Register(histogram)
	}

	return histogram
}

// NewConsumer create a Kafka Consumer with the chosen provider
// return: a new consumer or error.
func NewConsumer(kafkaURLs []string, topic, groupID string, provider types.Provider) (consumer.Consumer, error) {
	histogram := initPrometheus()

	switch provider {
	case types.Sarama:
		saramaConsumer, err := consumer.NewSaramaConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return &Subscriber{
			provider:  provider,
			consumer:  saramaConsumer,
			histogram: histogram,
		}, nil
	case types.Segmentio:
		segmentioConsumer, err := consumer.NewSegmentioConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return &Subscriber{
			provider:  provider,
			consumer:  segmentioConsumer,
			histogram: histogram,
		}, nil

	case types.Confluent:
		confluentConsumer, err := consumer.NewConfluentConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return &Subscriber{
			provider:  provider,
			consumer:  confluentConsumer,
			histogram: histogram,
		}, nil
	}

	return nil, fmt.Errorf("error: %w, message: %s, topic: %s, groupID: %s",
		ErrUndefinedProvider,
		"undefined provider or not informed",
		topic,
		groupID)
}

// Subscribe reads and return the message from the Topic.
func (s Subscriber) Subscribe(f func(message *message.Message) error) {
	defer s.trace(time.Now())
	s.consumer.Subscribe(f)
}

func (s Subscriber) trace(start time.Time) {
	seconds := time.Since(start).Seconds()
	prv := fmt.Sprintf("%v", s.provider)
	s.histogram.WithLabelValues(prv).Observe(seconds)
	log.WithFields(log.Fields{
		"provider": prv,
		"total":    seconds,
	}).Info("consumer seconds")
}

// Close finish consumer and stop read messages.
func (s Subscriber) Close() {
	s.consumer.Close()
}
