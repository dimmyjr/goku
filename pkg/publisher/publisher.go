package publisher

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dimmyjr/GoKafka/internal/producer"
	"github.com/dimmyjr/GoKafka/pkg/types"
	"github.com/prometheus/client_golang/prometheus"
)

var histogram *prometheus.HistogramVec

type publisher struct {
	provider types.Provider
	producer producer.Producer
}

func init() {
	histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "kafka_producer",
			Name:      "producer",
			Help:      "Producer execution in seconds",
			Buckets:   prometheus.DefBuckets,
		}, []string{"name"})

	prometheus.Register(histogram)
}

//NewProducer create a new producer
func NewProducer(kafkaURLs []string, topic string, provider *types.Provider) (producer.Producer, error) {
	switch *provider {
	case types.Sarama:
		saramaProducer, err := producer.NewSaramaProducer(kafkaURLs, topic)
		return publisher{
			provider: *provider,
			producer: saramaProducer,
		}, err
	case types.Segmentio:
		segmentioProducer, err := producer.NewSegmentioProducer(kafkaURLs, topic)
		return publisher{
			provider: *provider,
			producer: segmentioProducer,
		}, err

	case types.Confluent:
		confluentProducer, err := producer.NewConfluentProducer(kafkaURLs, topic)
		return publisher{
			provider: *provider,
			producer: confluentProducer,
		}, err
	}

	return nil, errors.New("invalid Provider to Producer")
}

func (pbs publisher) Publish(key, message string) error {
	start := time.Now()
	defer func() {
		time := time.Since(start).Seconds()
		prv := fmt.Sprintf("%v", pbs.provider)
		histogram.WithLabelValues(prv).Observe(time)
		log.WithFields(log.Fields{
			"provider": prv,
			"total":    time,
		}).Info("producer time")
	}()
	return pbs.producer.Publish(key, message)
}

func (pbs publisher) Close() {
	pbs.producer.Close()
}
