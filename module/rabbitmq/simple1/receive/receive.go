package main

import (
	"log"
	"rabbutmq/middleware"
	"time"

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
		"hello", // name
		false,   // durable: menyimpan data
		false,   // auto-delete: auto delete ketika tidak digunakan
		false,   // exclusive: memastikan konsumer yang diterima hanya satu
		false,   // no-wait
		nil,     // arguments
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

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	middleware.FailOnError(err, "Failed to set QoS")

// Qos nyala + autoack nyala = round robin
// Qos nyala + autoack mati = tidak round robin
// Qos mati + autoack nyala = round robin
// Qos mati + autoack mati = round robin

	// channel ini digunakan agar program tidak mati
	var forever chan string

	go func() {
		for d := range msgs {
			time.Sleep(time.Second * 4)
			log.Printf("Received a message: %s", d.Body)
			log.Printf("Routing Key: %s", d.RoutingKey)

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
