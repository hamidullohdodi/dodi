package service

import (
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

func (b *MsgBroker) CreateTransaction(body []byte) error {
	b.logger.Info("Publishing CreateTransaction message")
	return b.publishMessage("CreateTransaction", body)
}

func (b *MsgBroker) UpdateBudget(body []byte) error {
	b.logger.Info("Publishing UpdateBudget message")
	return b.publishMessage("UpdateBudget", body)
}

func (b *MsgBroker) UpdateGoal(body []byte) error {
	b.logger.Info("Publishing UpdateGoal message")
	return b.publishMessage("UpdateGoal", body)
}

func (b *MsgBroker) publishMessage(queueName string, body []byte) error {
	if b.channel == nil {
		b.logger.Error("Failed to publish message: channel is nil", "queue", queueName)
		return amqp.ErrClosed // Returning appropriate error
	}

	err := b.channel.Publish(
		"",        // exchange
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

	b.logger.Info("Message published", "queue", queueName)
	return nil
}
