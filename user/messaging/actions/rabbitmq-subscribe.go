package actions

import (
	"fmt"
	"g2/user/domain"
	"g2/user/variables"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitAddress string
)

type RabbitHandler interface {
	InsertUserAction(address, queueName, bindingKey string) error
	EditUserAction(address, queueName, bindingKey string) error
	RemoveUserAction(address, queueName, bindingKey string) error
}

type handler struct {
	domain  domain.ServiceInterface
	address string
}

func NewRabbitHandler(address string, service domain.ServiceInterface) RabbitHandler {
	rabbitAddress = address
	return &handler{domain: service, address: address}
}

func (c *handler) CreateChannel(address, queueName, bindingKey string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

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

func RabbitmqListenToActions(handler RabbitHandler) {
	go func() {
		err := handler.InsertUserAction(rabbitAddress, variables.InsertUserQueueName, variables.InsertUserBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.EditUserAction(rabbitAddress, variables.EditUserQueueName, variables.EditUserBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.RemoveUserAction(rabbitAddress, variables.RemoveUserQueueName, variables.RemoveUserBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
}
