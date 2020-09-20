package subscriber_test

import (
	"fmt"
	. "github.com/dimmyjr/GoKafka/internal/consumer"
	"github.com/dimmyjr/GoKafka/pkg/subscriber"
	"github.com/dimmyjr/GoKafka/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

const cmdProducer = "docker run --network host --rm confluentinc/cp-kafkacat " +
	"bash -c \"echo '%s' | kafkacat -b localhost:9092 -K: -P  -t %s\""

var _ = Describe("Subscriber", func() {

	rand.Seed(time.Now().UTC().UnixNano())

	Context("To produce message to kafka local", func() {
		kafkaURL := "localhost:9092"

		It("should produce message with Segmentio", func() {
			topic := strconv.Itoa(rand.Int()) + "-segmentio"
			message := "testes de kafka " + strconv.Itoa(rand.Int())
			provider, _ := types.GetProvider("segmentio")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string, 0)
			sub.Subscribe(func(input Message) error {
				c <- string(input.Value)
				return nil
			})
			Expect(<-c).Should(Equal(message))

		})

		It("should produce message with Sarama", func() {
			topic := strconv.Itoa(rand.Int()) + "-sarama"
			message := "testes de kafka " + strconv.Itoa(rand.Int())
			provider, _ := types.GetProvider("sarama")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string, 0)
			sub.Subscribe(func(input Message) error {
				c <- string(input.Value)
				return nil
			})
			Expect(<-c).Should(Equal(message))

		})

		It("should produce message with Confluent", func() {
			topic := strconv.Itoa(rand.Int()) + "-confluent"
			message := "testes de kafka " + strconv.Itoa(rand.Int())
			provider, _ := types.GetProvider("confluent")

			By("creating sub")
			sub, err := subscriber.NewConsumer([]string{kafkaURL}, topic, "test", provider)
			Expect(err).NotTo(HaveOccurred())

			By("publishing message")
			_, err = producer(topic, message)
			Expect(err).NotTo(HaveOccurred())

			By("should get message produced")

			c := make(chan string, 0)
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
