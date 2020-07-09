package main

import (
	_ "expvar"
	"flag"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	server "github.com/dimmyjr/GoKafka/cmd"
)

var kafkaURL, topic, provider, grpcPort, prometheusPort string

func init() {
	flag.StringVar(&kafkaURL, "kafkaURL", "", "Kafka Url")
	flag.StringVar(&topic, "topic", "", "Topic Name")
	flag.StringVar(&provider, "provider", "", "confluent, sarama, segmentio")
	flag.StringVar(&grpcPort, "grpcPort", "", "gRPC Port")
	flag.StringVar(&prometheusPort, "prometheusPort", "", "Prometheus Port")
	flag.Parse()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+prometheusPort, nil)
	server.Start(kafkaURL, topic, provider, grpcPort)
}
