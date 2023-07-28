package actions

import (
	"encoding/json"
	"fmt"
	"g2/todo/domain"
	"g2/todo/variables"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *handler) InsertTodoListAction(address, queueName, bindingKey string) error {
	ch, err := c.CreateChannel(address, queueName, bindingKey)
	if err != nil {
		return err
	}
	defer ch.Close()
	deliveries, err := ch.Consume(
		queueName,
		variables.ExchangeName,
		variables.ConsumeAutoAck,
		variables.ConsumeExclusive,
		variables.ConsumeNoLocal,
		variables.ConsumeNoWait,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range deliveries {
		//print the message to the console
		fmt.Println("i get a message")
		go func() {
			message := &domain.TodoLists{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.InsertTodoListService(message)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
func (c *handler) EditTodoListAction(address, queueName, bindingKey string) error {
	ch, err := c.CreateChannel(address, queueName, bindingKey)
	if err != nil {
		return err
	}
	defer ch.Close()
	deliveries, err := ch.Consume(
		queueName,
		variables.ExchangeName,
		variables.ConsumeAutoAck,
		variables.ConsumeExclusive,
		variables.ConsumeNoLocal,
		variables.ConsumeNoWait,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range deliveries {
		//print the message to the console
		fmt.Println("i get a message")
		go func() {
			message := &domain.TodoLists{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.EditTodoListService(message)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
func (c *handler) RemoveTodoListAction(address, queueName, bindingKey string) error {
	ch, err := c.CreateChannel(address, queueName, bindingKey)
	if err != nil {
		return err
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		queueName,
		variables.ExchangeName,
		variables.ConsumeAutoAck,
		variables.ConsumeExclusive,
		variables.ConsumeNoLocal,
		variables.ConsumeNoWait,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range deliveries {
		//print the message to the console
		fmt.Println("i get a message")
		go func() {
			message := &domain.TodoLists{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.RemoveTodoListService(message.Id, message.UserID)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
