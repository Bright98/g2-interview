package actions

import (
	"encoding/json"
	"fmt"
	"g2/user/domain"
	"g2/user/variables"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *handler) InsertUserAction(address, queueName, bindingKey string) error {
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
			message := &domain.Users{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.InsertUserService(message)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
func (c *handler) EditUserAction(address, queueName, bindingKey string) error {
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
			message := &domain.Users{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.EditUserService(message)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
func (c *handler) RemoveUserAction(address, queueName, bindingKey string) error {
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
			message := &domain.Users{}
			err = json.Unmarshal(msg.Body, message)
			if err != nil {
				fmt.Print(err)
			}
			c.domain.RemoveUserService(message.Id)
		}()
		// Acknowledge that we have received the message. it can be removed from the queue
		msg.Ack(false)
	}
	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	return chanErr
}
