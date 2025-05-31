package main

import (
	"fmt"

	"github.com/felipeazsantos/pos-goexpert/events/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	err := rabbitmq.Publish(ch, "Hello, RabbitMQ!")
	if err != nil {
		fmt.Printf("Error publishing message: %v\n", err)
		return
	}
}