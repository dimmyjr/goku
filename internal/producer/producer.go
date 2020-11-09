package producer

// Producer to publish to kafka.
type Producer interface {
	Publish(key, msg string) error
	Close()
}
