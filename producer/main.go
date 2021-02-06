package producer

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Publish(msg string, topic string) error {
	message := {
		TopicPartition: totopic,
		Value: []byte(msg),
	}
	return nil
}