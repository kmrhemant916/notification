package main

import (
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
}