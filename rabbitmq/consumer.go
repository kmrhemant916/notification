package rabbitmq

import (
	"log"

	"github.com/kmrhemant916/notification/utils"
	"github.com/rabbitmq/amqp091-go"
)


func (r *Rabbitmq) DeclareConsumer(q amqp091.Queue, ch *amqp091.Channel) {
	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	utils.FailOnError(err, "Failed to register a consumer")
	// Start consuming messages
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()
	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}