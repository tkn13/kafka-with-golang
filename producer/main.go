package main

import (
	"encoding/json"
	"fmt"
	"model"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	count := 1
	for {
		count += 1
		msg := model.TestMessage{
			Id:   count,
			Name: "test",
			Time: time.Now(),
		}

		msgByte, _ := json.Marshal(msg)

		kafkaMsg := sarama.ProducerMessage{
			Topic: "testTopic",
			Value: sarama.StringEncoder(msgByte),
		}

		p, o, err := producer.SendMessage(&kafkaMsg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message was sent to partition: %v, offset: %v\n", p, o)
		time.Sleep(3 * time.Second)
	}
}
