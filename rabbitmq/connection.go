package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	Config = "config/config.yaml"
)

func (r *Rabbitmq) Connection() (*amqp.Connection, error) {
    connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
        r.Username,
        r.Password,
        r.Host,
        r.Port,
    )
    conn, err := amqp.Dial(connectionString)
    if err != nil {
        return nil, err
    }
    return conn, nil
}
