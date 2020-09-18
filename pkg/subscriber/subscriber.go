package subscriber

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dimmyjr/GoKafka/internal/consumer"
	"github.com/dimmyjr/GoKafka/pkg/types"
	"github.com/prometheus/client_golang/prometheus"
)

var histogram *prometheus.HistogramVec

type subscriber struct {
	provider types.Provider
	consumer consumer.Consumer
}

func init() {
	histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "kafka_consumer",
			Name:      "consumer",
			Help:      "Consumer execution in seconds",
			Buckets:   prometheus.DefBuckets,
		}, []string{"name"})

	prometheus.Register(histogram)
}

//NewConsumer create a new consumer
func NewConsumer(kafkaURLs []string, topic, groupID string, provider *types.Provider) (consumer.Consumer, error) {
	switch *provider {
	case types.Sarama:
		saramaConsumer, err := consumer.NewSaramaConsumer(kafkaURLs, topic, groupID)
		return subscriber{
			provider: *provider,
			consumer: saramaConsumer,
		}, err
	case types.Segmentio:
		segmentioConsumer, err := consumer.NewSegmentioConsumer(kafkaURLs, topic, groupID)
		return subscriber{
			provider: *provider,
			consumer: segmentioConsumer,
		}, err

	case types.Confluent:
		confluentConsumer, err := consumer.NewConfluentConsumer(kafkaURLs, topic, groupID)
		return subscriber{
			provider: *provider,
			consumer: confluentConsumer,
		}, err
	}

	return nil, errors.New("invalid Provider to Consumer")
}

func (pbs subscriber) Subscribe(f func(message consumer.Message) error) {
	start := time.Now()
	defer func() {
		time := time.Since(start).Seconds()
		prv := fmt.Sprintf("%v", pbs.provider)
		histogram.WithLabelValues(prv).Observe(time)
		log.WithFields(log.Fields{
			"provider": prv,
			"total":    time,
		}).Info("consumer time")
	}()

	pbs.consumer.Subscribe(f)
}

func (pbs subscriber) Close() {
	pbs.consumer.Close()
}
