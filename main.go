package main

import (
	"log"

	"github.com/kmrhemant916/notification/rabbitmq"
	"github.com/kmrhemant916/notification/utils"
)

const (
	Config = "config/config.yaml"
)

func main() {
	var config utils.Config
	c, err:= config.ReadConf(Config)
    if err != nil {
        panic(err)
    }
	rabbitmqConfig := c.Rabbitmq
	conn, err := rabbitmq.Connection(rabbitmqConfig.Username, rabbitmqConfig.Password, rabbitmqConfig.Host, rabbitmqConfig.Port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"mail", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}
	var forever chan struct{}

	go func() {
	  for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	  }
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}