package subscriber_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os/exec"

	"github.com/dimmyjr/goku/message"
	"github.com/dimmyjr/goku/subscriber"
	"github.com/dimmyjr/goku/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const cmdProducer = "docker run --network host --rm confluentinc/cp-kafkacat " +
	"bash -c \"echo '%s' | kafkacat -b localhost:9092 -K: -P  -t %s\""

var _ = Describe("Subscriber", func() {
	Context("To produce message to kafka local", func() {
		kafkaURL := "localhost:9092"

		It("should produce message with Segmentio", func() { //nolint:dupl
			topic := randomString() + "-segmentio"
			msg := "testes de kafka " + randomString()
			provider := types.Segmentio

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, msg)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input *message.Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(msg))
		})

		It("should produce message with Sarama", func() { //nolint:dupl
			topic := randomString() + "-sarama"
			msg := "testes de kafka " + randomString()
			provider := types.Sarama

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, msg)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input *message.Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(msg))
		})

		It("should produce message with Confluent", func() { //nolint:dupl
			topic := randomString() + "-confluent"
			msg := "testes de kafka " + randomString()
			provider := types.Confluent

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, msg)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input *message.Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(msg))
		})
	})
})

func producer(topic, message string) ([]byte, error) {
	cmd := fmt.Sprintf(cmdProducer, "1:"+message, topic)

	return exec.Command("/bin/sh", "-c", cmd).Output()
}

func randomString() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))

	return n.String()
}
