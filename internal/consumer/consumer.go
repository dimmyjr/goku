package consumer

import (
	"github.com/dimmyjr/goku/message"
)

// Consumer kafka.
type Consumer interface {
	Subscribe(f func(message *message.Message) error)
	Close()
}
