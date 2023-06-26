package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("error", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("error", err)
	}

	err = ch.ExchangeDeclare(
		"test_topic",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("error", err)
	}

	q, err := ch.QueueDeclare(
		"q1",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("error", err)
	}

	err = ch.QueueBind(
		q.Name,
		"key_topic",
		"test_topic",
		false,
		nil,
	)
	if err != nil {
		fmt.Println("error", err)
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("error", err)
	}

	var stop chan struct{}

	go func() {
		for v := range msg {
			log.Println("hasilnya ", string(v.Body))

			v.Ack(true)
		}

	}()

	<-stop

}
