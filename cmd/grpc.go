package server

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/dimmyjr/GoKafka/internal/producer"
	"github.com/dimmyjr/GoKafka/pkg/publisher"
	types "github.com/dimmyjr/GoKafka/pkg/types"
	uuid "github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	producer producer.Producer
}

func (s *server) Publish(ctx context.Context, in *Request) (*Reply, error) {
	go s.producer.Publish(uuid.New().String(), in.GetMessage())
	return &Reply{Message: " " + in.GetMessage()}, nil
}

// Start start gRPC Server with Kafka connection
func Start(kafkaURL, topic, provider, grpcPort string) {
	prv, err := types.GetProvider(provider)
	if err != nil {
		log.Fatalf("Provider invalid: %v", err)
	}

	prdc, err := publisher.NewProducer([]string{kafkaURL}, topic+"-sarama", prv)
	defer prdc.Close()

	if err != nil {
		log.Fatalf("Error create producer: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterProducerServer(s, &server{producer: producer})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
