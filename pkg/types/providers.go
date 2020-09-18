package types

import "errors"

// Provider enum
type Provider string

const (
	// Sarama provider
	Sarama Provider = "sarama"
	// Segmentio provider
	Segmentio Provider = "segmentio"
	// Confluent provider
	Confluent Provider = "confluent"
)

func GetProvider(input string) (*Provider, error) {
	p := Provider(input)
	switch p {
	case Sarama, Segmentio, Confluent:
		return &p, nil
	default:
		return nil, errors.New("invalid leave type")
	}

}
