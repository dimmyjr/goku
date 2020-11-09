package consumer

import (
	"reflect"
	"time"
)

// Consumer kafka.
type Consumer interface {
	Subscribe(f func(message Message) error)
	Close()
}

// Message to read/write.
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

// Header from message.
type Header struct {
	Key   string
	Value []byte
}

func header(headers interface{}) []Header {
	var result []Header

	values := reflect.ValueOf(headers)

	for i := 0; i < values.Len(); i++ {
		h := values.Index(i)
		key := reflect.ValueOf(h).Elem().FieldByName("key").String()
		value := reflect.ValueOf(h).Elem().FieldByName("value").Bytes()

		result = append(result, Header{
			Key:   key,
			Value: value,
		})
	}

	return result
}
