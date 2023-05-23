package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sales-go/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type publisher struct {
}

func NewPublisher() *publisher {
	return &publisher{}
}

type PublisherInterface interface {
	Publish(body interface{}) (err error)
}

func (p *publisher) Publish(body interface{}) (err error) {
	config, err := config.LoadConfig()
	if err != nil {
		err = fmt.Errorf("failed to load config")
		return
	}

	conn, err := amqp.Dial(config.RabbitMQURL)
	if err != nil {
		err = fmt.Errorf("failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		err = fmt.Errorf("failed to open a channel")
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"create_transaction", // name
		true,         // durable
		false,        // auto delete queue when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		err = fmt.Errorf("failed to declare a queue")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonByte, err := json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("error marshal")
		return
	}

	err = ch.PublishWithContext(ctx,
		"",     	// exchange
		q.Name, 	// routing key
		false,  	// mandatory
		false,		// immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         jsonByte,
		})
	if err != nil {
		err = fmt.Errorf("failed to publish a message")
		return
	}

	log.Printf(" [x] Sent %s", body)
	return
}