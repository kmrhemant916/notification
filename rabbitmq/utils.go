package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func (r *Rabbitmq) DeclareChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	return ch, err
}

func (r *Rabbitmq) DeclareExchange(ch *amqp.Channel) (error) {
	return ch.ExchangeDeclare(
		r.Exchange.Name, // name
		r.Exchange.Kind,      // kind
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (r *Rabbitmq) DeclareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return q, err
}

func (r *Rabbitmq) DeclareQueueBind(ch *amqp.Channel, q amqp.Queue, k string) (error) {
	return ch.QueueBind(
		q.Name,      // queue name
		k,  // routing key
		r.Exchange.Name, // exchange name
		false,       // no-wait
		nil,         // arguments
	)
}