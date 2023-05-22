package main

import (
	"context"
	"log"
	"os"
	"rabbutmq/middleware"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// menyambungkan ke koneksi rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	middleware.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   //name of exchange
		"fanout", //type there are several type of exchange, direct, topic, header and fanout
		true,     //durable
		false,
		false,
		false,
		nil,
	)
	middleware.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(
		ctx,
		"",    // exchange
		"",    // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	middleware.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
