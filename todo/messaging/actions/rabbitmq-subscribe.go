package actions

import (
	"fmt"
	"g2/todo/domain"
	"g2/todo/variables"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	rabbitAddress string
)

type RabbitHandler interface {
	InsertTodoListAction(address, queueName, bindingKey string) error
	EditTodoListAction(address, queueName, bindingKey string) error
	RemoveTodoListAction(address, queueName, bindingKey string) error
	InsertTodoItemAction(address, queueName, bindingKey string) error
	EditTodoItemAction(address, queueName, bindingKey string) error
	RemoveTodoItemAction(address, queueName, bindingKey string) error
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
	//todo list
	go func() {
		err := handler.InsertTodoListAction(rabbitAddress, variables.InsertTodoListQueueName, variables.InsertTodoListBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.EditTodoListAction(rabbitAddress, variables.EditTodoListQueueName, variables.EditTodoListBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.RemoveTodoListAction(rabbitAddress, variables.RemoveTodoListQueueName, variables.RemoveTodoListBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()

	//todo item
	go func() {
		err := handler.InsertTodoItemAction(rabbitAddress, variables.InsertTodoItemQueueName, variables.InsertTodoItemBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.EditTodoItemAction(rabbitAddress, variables.EditTodoItemQueueName, variables.EditTodoItemBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
	go func() {
		err := handler.RemoveTodoItemAction(rabbitAddress, variables.RemoveTodoItemQueueName, variables.RemoveTodoItemBindingKey)
		if err != nil {
			fmt.Println("rabbit error: ", err.Error())
		}
	}()
}
