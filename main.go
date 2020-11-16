package main

import (
	_ "expvar"
	"flag"
	"log"
	"net/http"

	"github.com/dimmyjr/goku/cmd/server"
	"github.com/dimmyjr/goku/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var kafkaURL, topic, provider, grpcPort, prometheusPort string

	flag.StringVar(&kafkaURL, "kafkaURL", "", "Kafka Url")
	flag.StringVar(&topic, "topic", "", "Topic Name")
	flag.StringVar(&provider, "provider", "", "confluent, sarama, segmentio")
	flag.StringVar(&grpcPort, "grpcPort", "", "gRPC Port")
	flag.StringVar(&prometheusPort, "prometheusPort", "", "Prometheus Port")
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())

	go prometheus(prometheusPort)
	server.Start(kafkaURL, topic, grpcPort, types.Provider(provider))
}

func prometheus(prometheusPort string) {
	err := http.ListenAndServe(":"+prometheusPort, nil)
	if err != nil {
		log.Fatal("Error to initialize prometheus metrics")
	}
}
