package events

import (
	"g2/api-gateway/variables"
	"github.com/streadway/amqp"
	"log"
)

type RabbitConsumer struct {
	amqpConn *amqp.Connection
	address  string
}

func NewRabbitMQ(address string) *RabbitConsumer {
	conn, err := amqp.Dial(address)
	if err != nil {
		log.Fatalln(err)
	}

	rabbit := &RabbitConsumer{}
	rabbit.amqpConn = conn
	rabbit.address = address
	return rabbit
}

// CreateChannel Consume messages
func (c *RabbitConsumer) CreateChannel(queueName, bindingKey string) (*amqp.Channel, error) {
	conn := c.amqpConn
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		variables.ExchangeName,
		variables.ExchangeKind,
		variables.ExchangeDurable,
		variables.ExchangeAutoDelete,
		variables.ExchangeInternal,
		variables.ExchangeNoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}
	queue, err := ch.QueueDeclare(
		queueName,
		variables.QueueDurable,
		variables.QueueAutoDelete,
		variables.QueueExclusive,
		variables.QueueNoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		variables.ExchangeName,
		variables.QueueNoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.Qos(
		variables.PrefetchCount,  // prefetch count
		variables.PrefetchSize,   // prefetch size
		variables.PrefetchGlobal, // global
	)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
func (c *RabbitConsumer) PublishMessage(body []byte, queueName, bindingKey string) error {
	ch, err := c.CreateChannel(queueName, bindingKey)
	if err != nil {
		return err
	}
	message := amqp.Publishing{
		Body: body,
	}
	err = ch.Publish(variables.ExchangeName, bindingKey, false, false, message)
	if err != nil {
		return err
	}
	return nil
}
