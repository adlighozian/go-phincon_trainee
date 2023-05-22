package db

import (
	"inventory/middleware"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitConn() *amqp.Channel {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	ch, err := conn.Channel()
	middleware.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	return ch
}
