package main

import (
	"consumer/constant"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {

	server := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(server, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(constant.MAIN_TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	fmt.Print("Consumer started\n")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Print(err)

		case msg := <-partitionConsumer.Messages():
			fmt.Printf("%v\n", string(msg.Value))
		}
	}
}
