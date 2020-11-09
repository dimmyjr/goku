package types

import (
	"errors"
	"fmt"
)

// Provider enum.
type Provider string

var ErrUndefinedProvider = errors.New("undefinedProvider")

const (
	// Sarama provider.
	Sarama Provider = "sarama"
	// Segmentio provider.
	Segmentio Provider = "segmentio"
	// Confluent provider.
	Confluent Provider = "confluent"
)

func GetProvider(input string) (Provider, error) {
	p := Provider(input)
	switch p {
	case Sarama, Segmentio, Confluent:
		return p, nil
	default:
		return "", fmt.Errorf("error: %w", ErrUndefinedProvider)
	}
}
