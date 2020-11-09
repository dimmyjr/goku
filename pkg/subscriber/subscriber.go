package subscriber

import (
	"errors"
	"fmt"
	"time"

	"github.com/dimmyjr/goku/internal/consumer"
	"github.com/dimmyjr/goku/pkg/types"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var ErrUndefinedProvider = errors.New("undefinedProvider")

type subscriber struct {
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

// NewConsumer create a Kafka Consumer with the chosen provider.
// @topic[required]: existing topic in the Kafka cluster
// @groupID: is to load balance the produced data
//			(if the groupID is different for each consumer, each consumer will get the copy of data)
// @provider: Sarama, Segmentio, Confluent
// return: a new consumer or error.
func NewConsumer(kafkaURLs []string, topic, groupID string, provider types.Provider) (consumer.Consumer, error) {
	histogram := initPrometheus()

	switch provider {
	case types.Sarama:
		saramaConsumer, err := consumer.NewSaramaConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return subscriber{
			provider:  provider,
			consumer:  saramaConsumer,
			histogram: histogram,
		}, nil
	case types.Segmentio:
		segmentioConsumer, err := consumer.NewSegmentioConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return subscriber{
			provider:  provider,
			consumer:  segmentioConsumer,
			histogram: histogram,
		}, nil

	case types.Confluent:
		confluentConsumer, err := consumer.NewConfluentConsumer(kafkaURLs, topic, groupID)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s, groupID: %s", err, topic, groupID)
		}

		return subscriber{
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
func (service subscriber) Subscribe(f func(message consumer.Message) error) {
	defer service.trace(time.Now())
	service.consumer.Subscribe(f)
}

func (service subscriber) trace(start time.Time) {
	seconds := time.Since(start).Seconds()
	prv := fmt.Sprintf("%v", service.provider)
	service.histogram.WithLabelValues(prv).Observe(seconds)
	log.WithFields(log.Fields{
		"provider": prv,
		"total":    seconds,
	}).Info("consumer seconds")
}

// Close finish consumer and stop read messages.
func (service subscriber) Close() {
	service.consumer.Close()
}
