package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/dimmyjr/goku/internal/producer"
	"github.com/dimmyjr/goku/pkg/publisher"
	"github.com/dimmyjr/goku/pkg/types"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var ErrInvalidProducer = errors.New("invalidProducer")

type server struct {
	producer producer.Producer
}

// Publish to produce a message to Kafka on gRPC.
func (s *server) Publish(ctx context.Context, in *Request) (*Reply, error) {
	if s.producer == nil {
		return nil, fmt.Errorf("error: %w", ErrInvalidProducer)
	}

	_ = s.producer.Publish(uuid.New().String(), in.GetMessage())

	return &Reply{Message: " " + in.GetMessage()}, nil
}

// Start gRPC Server with Kafka connection.
func Start(kafkaURL, topic, provider, grpcPort string) {
	prv, err := types.GetProvider(provider)
	if err != nil {
		log.Fatalf("Provider invalid: %v", err)
	}

	prdc, err := publisher.NewProducer([]string{kafkaURL}, topic+"-sarama", prv)
	if err != nil {
		log.Fatalf("Error create producer: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterProducerServer(s, &server{producer: prdc})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer prdc.Close()
}
