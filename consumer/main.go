package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer() error{
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id": "consumergroup",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("received: ", string(msg.Value))
		}
	}
}