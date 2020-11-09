package publisher

import (
	"errors"
	"fmt"
	"time"

	"github.com/dimmyjr/GoKafka/internal/producer"
	"github.com/dimmyjr/GoKafka/pkg/types"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var ErrUndefinedProvider = errors.New("undefinedProvider")

type publisher struct {
	histogram *prometheus.HistogramVec
	provider  types.Provider
	producer  producer.Producer
}

func initPrometheus() *prometheus.HistogramVec {
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem:   "Publisher",
			ConstLabels: prometheus.Labels{},
			Namespace:   "kafka_producer",
			Name:        "Producer",
			Help:        "Producer execution in seconds",
			Buckets:     prometheus.DefBuckets,
		}, []string{"name"})

	if histogram != nil {
		_ = prometheus.Register(histogram)
	}

	return histogram
}

// NewProducer create a Kafka producer with the chosen provider
// @topic[required]: existing topic in the Kafka cluster
// @provider: Sarama, Segmentio, Confluent
// return: a new producer or error.
func NewProducer(kafkaURLs []string, topic string, provider types.Provider) (producer.Producer, error) {
	histogram := initPrometheus()

	switch provider {
	case types.Sarama:
		saramaProducer, err := producer.NewSaramaProducer(kafkaURLs, topic)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s", err, topic)
		}

		return publisher{
			provider:  provider,
			producer:  saramaProducer,
			histogram: histogram,
		}, nil
	case types.Segmentio:
		segmentioProducer, err := producer.NewSegmentioProducer(kafkaURLs, topic)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s", err, topic)
		}

		return publisher{
			provider:  provider,
			producer:  segmentioProducer,
			histogram: histogram,
		}, nil

	case types.Confluent:
		confluentProducer, err := producer.NewConfluentProducer(kafkaURLs, topic)
		if err != nil {
			return nil, fmt.Errorf("error: %w, topic: %s", err, topic)
		}

		return publisher{
			provider:  provider,
			producer:  confluentProducer,
			histogram: histogram,
		}, nil
	}

	return nil, fmt.Errorf("error: %w, topic: %s", ErrUndefinedProvider, topic)
}

func (service publisher) Publish(key, message string) error {
	defer service.trace(time.Now())

	return service.producer.Publish(key, message)
}

func (service publisher) Close() {
	service.producer.Close()
}

func (service publisher) trace(start time.Time) {
	seconds := time.Since(start).Seconds()
	prv := fmt.Sprintf("%v", service.provider)
	service.histogram.WithLabelValues(prv).Observe(seconds)
	log.WithFields(log.Fields{
		"provider": prv,
		"total":    seconds,
	}).Info("producer seconds")
}
