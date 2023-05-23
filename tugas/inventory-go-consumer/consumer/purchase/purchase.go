package main

import (
	"encoding/json"
	"inventory/middleware"
	"inventory/model"
	"inventory/repository"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//membuat koneksi
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	middleware.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//membuat channel
	ch, err := conn.Channel()
	middleware.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// membuat queue
	q, err := ch.QueueDeclare(
		"Purchase1", // name
		false,       // durable: menyimpan data
		false,       // auto-delete: auto delete ketika tidak digunakan
		false,       // exclusive: memastikan konsumer yang diterima hanya satu
		false,       // no-wait
		nil,         // arguments
	)
	middleware.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack: untuk memberitahu ke rabbit kalau pesan sudah dikirim, lalu rabbit akan menghapus pesannya jika false
		false,  // exclusive:
		false,  // no-local:
		false,  // no-wait:
		nil,    // args
	)
	middleware.FailOnError(err, "Failed to register a consumer")

	var forever chan string

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// log.Printf("Routing Key: %s", d.RoutingKy)
			var data []model.SendPurchase
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Println("Error deserializing message:", err)
				continue
			}
			d.Ack(false)
			repository.TambahPurchase(data)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
