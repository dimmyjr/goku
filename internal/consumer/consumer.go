package consumer

import (
	"time"
)

// Consumer messages from kafka
type Consumer interface {
	Subscribe(f func(message Message) error)
	Close()
}

type Message struct {
	// Topic is reads only and MUST NOT be set when writing messages
	Topic string

	// Partition is reads only and MUST NOT be set when writing messages
	Partition int32
	Offset    int64
	Key       []byte
	Value     []byte
	Headers   []Header

	// If not set at the creation, Time will be automatically set when
	// writing the message.
	Time time.Time
}

type Header struct {
	Key   string
	Value []byte
}
