package service

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type MsgBroker struct {
	channel *amqp.Channel
	logger  *slog.Logger
}

func NewMsgBroker(channel *amqp.Channel, logger *slog.Logger) *MsgBroker {
	if channel == nil {
		logger.Error("Failed to initialize MsgBroker: channel is nil")
		return nil
	}
	return &MsgBroker{
		channel: channel,
		logger:  logger,
	}
}

func (b *MsgBroker) PostComment(body []byte) error {
	b.logger.Info("Publishing PostComment message")
	return b.publishMessage("PostComment", body)
}

func (b *MsgBroker) UpdateComment(body []byte) error {
	b.logger.Info("Publishing UpdateComment message")
	return b.publishMessage("UpdateComment", body)
}

func (b *MsgBroker) AddLike(body []byte) error {
	b.logger.Info("Publishing AddLike message")
	return b.publishMessage("AddLike", body)
}

func (b *MsgBroker) PostTweet(body []byte) error {
	b.logger.Info("Publishing PostTweet message")
	return b.publishMessage("PostTweet", body)
}

func (b *MsgBroker) UpdateTweet(body []byte) error {
	b.logger.Info("Publishing UpdateTweet message")
	return b.publishMessage("UpdateTweet", body)
}

func (b *MsgBroker) publishMessage(queueName string, body []byte) error {
	if b.channel == nil {
		b.logger.Error("Failed to publish message: channel is nil", "queue", queueName)
		return errors.New("failed to publish message: channel is nil") // Creating a more descriptive error
	}

	err := b.channel.Publish(
		"",        // exchange, keeping it blank if default exchange is intended
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		b.logger.Error("Failed to publish message", "queue", queueName, "error", err.Error())
		return err
	}

	b.logger.Info("Message published successfully", "queue", queueName)
	return nil
}
