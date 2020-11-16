package server

import (
	"context"
	"log"
	"net"

	"github.com/dimmyjr/goku/publisher"
	"github.com/dimmyjr/goku/types"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	publisher *publisher.Publisher
}

// Publish to produce a message to Kafka on gRPC.
func (s *server) Publish(ctx context.Context, in *Request) (*Reply, error) {
	_ = s.publisher.Publish(uuid.New().String(), in.GetMessage())

	return &Reply{Message: " " + in.GetMessage()}, nil
}

// Start gRPC Server with Kafka connection.
func Start(kafkaURL, topic, grpcPort string, provider types.Provider) {
	produce, err := publisher.New([]string{kafkaURL}, topic+"-sarama", provider)
	if err != nil {
		log.Fatalf("Error create producer: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterProducerServer(s, &server{publisher: produce})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer produce.Close()
}
