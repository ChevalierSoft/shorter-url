package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func setProducer() *sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"http://redpanda:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		panic("failled to create producer")
	}
	defer producer.Close()

	return &producer
}
