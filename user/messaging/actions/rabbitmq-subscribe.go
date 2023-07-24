package actions

import (
	"fmt"
	"g2/user/domain"
	"g2/user/variables"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

type RabbitConsumer struct {
	amqpConn *amqp.Connection
}
type RabbitHandler interface {
	InsertUserAction(exchange, queueName, bindingKey, consumerTag string) error
	EditUserAction(exchange, queueName, bindingKey, consumerTag string) error
	RemoveUserAction(exchange, queueName, bindingKey, consumerTag string) error
}

type handler struct {
	domain domain.ServiceInterface
}

func NewRabbitHandler(service domain.ServiceInterface) RabbitHandler {
	return &handler{domain: service}
}

func (c *handler) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		exchangeName,
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
		exchangeName,
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

func RabbitmqListenToActions(handler RabbitHandler) {
	go func() {
		err := handler.InsertUserAction(
			variables.ExchangeName,
			variables.InsertUserQueueName,
			variables.InsertUserQueueName,
			"",
		)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.RemoveUserAction(
			variables.ExchangeName,
			variables.RemoveUserQueueName,
			variables.RemoveUserQueueName,
			"",
		)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.RemoveUserAction(
			variables.ExchangeName,
			variables.RemoveUserQueueName,
			variables.RemoveUserQueueName,
			"",
		)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
}
