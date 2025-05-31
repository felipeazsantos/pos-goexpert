package main

import (
	"fmt"

	"github.com/felipeazsantos/pos-goexpert/events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch,  msgs)

	for msg := range msgs {
		fmt.Printf("Received message: %s\n", string(msg.Body))
		if err := msg.Ack(false); err != nil {
			fmt.Println("Error acknowledging message:", err)
		}
	}

}
