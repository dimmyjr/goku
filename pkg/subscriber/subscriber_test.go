package subscriber_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os/exec"

	. "github.com/dimmyjr/goku/internal/consumer"
	"github.com/dimmyjr/goku/pkg/subscriber"
	"github.com/dimmyjr/goku/pkg/types"
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
			message := "testes de kafka " + randomString()
			provider, _ := types.GetProvider("segmentio")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(message))
		})

		It("should produce message with Sarama", func() { //nolint:dupl
			topic := randomString() + "-sarama"
			message := "testes de kafka " + randomString()
			provider, _ := types.GetProvider("sarama")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(message))
		})

		It("should produce message with Confluent", func() { //nolint:dupl
			topic := randomString() + "-confluent"
			message := "testes de kafka " + randomString()
			provider, _ := types.GetProvider("confluent")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string)
			sub.Subscribe(func(input Message) error {
				c <- string(input.Value)

				return nil
			})
			Expect(<-c).Should(Equal(message))
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
