package publisher

import (
	"context"
	"encoding/json"
	"inventory/helper/middleware"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type purchase struct {
}

func Newpurchase() *purchase {
	return &purchase{}
}

type PurchaseInterface interface {
	PubPurchase(body interface{}) error
}

func (p *purchase) PubPurchase(data interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// membuat channel
	ch, err := conn.Channel()
	middleware.FailError(err, "Failed to open a channel")
	defer ch.Close()

	// membuat queue
	q, err := ch.QueueDeclare(
		"Purchase1", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	middleware.FailError(err, "Failed to declare a queue")

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.PublishWithContext(
		ctx,    // context
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	middleware.FailError(err, "Failed to publish a message")
	log.Printf(" [x] Sent")

	return nil
}
