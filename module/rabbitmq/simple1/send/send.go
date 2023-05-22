package main

import (
	"context"
	"log"
	"rabbutmq/middleware"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	for i := 0; i < 5; i++ {
		send()
	}

}

func send() {
	// menyambungkan ke koneksi rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// membuat channel
	ch, err := conn.Channel()
	middleware.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// membuat queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	middleware.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(
		ctx,    // context
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	middleware.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
