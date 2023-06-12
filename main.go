package main

import (
	"github.com/kmrhemant916/notification/rabbitmq"
	"github.com/kmrhemant916/notification/utils"
)

const (
	Config = "config/config.yaml"
	BindingKey = "*.mail"
)

func main() {
	rabbit := rabbitmq.Setup()
	conn, err := rabbit.Connection()
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := rabbit.DeclareChannel(conn)
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	err = rabbit.DeclareExchange(ch)
	utils.FailOnError(err, "Failed to declare an exchange")
	q, err := rabbit.DeclareQueue(ch)
	utils.FailOnError(err, "Failed to declare a queue")
	err = rabbit.DeclareQueueBind(ch, q, BindingKey, )
	utils.FailOnError(err, "Failed to bind a queue")
	rabbit.DeclareConsumer(q, ch)
}