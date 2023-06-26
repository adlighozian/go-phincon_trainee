package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("error 1", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("error 2", err)
	}

	errs := ch.ExchangeDeclare(
		"test_topic",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if errs != nil {
		log.Println("error 3", errs)
	}

	sent := "test"

	ch.PublishWithContext(ctx,
		"test_topic",
		"key_topic",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(sent),
		},
	)

}
