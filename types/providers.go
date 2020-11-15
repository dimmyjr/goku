// Supported types
package types

import (
	"errors"
)

// Provider enum.
type Provider string

// ErrUndefinedProvider : undefined provider Error.
var ErrUndefinedProvider = errors.New("undefinedProvider")

const (
	// Sarama provider.
	Sarama Provider = "sarama"
	// Segmentio provider.
	Segmentio Provider = "segmentio"
	// Confluent provider.
	Confluent Provider = "confluent"
)

func (p Provider) String() string {
	return string(p)
}
