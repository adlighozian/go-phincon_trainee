package main

import (
	"log"
	"rabbutmq/middleware"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	middleware.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"query", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	middleware.FailOnError(err, "Failed to declare a queue")

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

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	middleware.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	middleware.FailOnError(err, "Failed to register a consumer")

	var forever chan string
	var counter int = 0

	go func() {
		for d := range msgs {
			counter++
			log.Println("Received a message: ", string(d.Body))
			log.Println("Received content/type : ", d.RoutingKey)
			log.Println(counter)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
