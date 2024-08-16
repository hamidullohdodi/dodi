package producer

import (
	"log/slog"

	rmq "github.com/rabbitmq/amqp091-go"
)

type (
	MsgBroker struct {
		channel *rmq.Channel
		logger  *slog.Logger
	}
)

func NewMsgBroker(channel *rmq.Channel, logger *slog.Logger) *MsgBroker {
	return &MsgBroker{
		channel: channel,
		logger:  logger,
	}
}

func (b *MsgBroker) CreateBooking(body []byte) error {
	return b.publishMessage("create_booking", body)
}
func (b *MsgBroker) CancelBooking(body []byte) error {
	return b.publishMessage("booking_cancelled", body)
}
func (b *MsgBroker) Payment(body []byte) error {
	return b.publishMessage("payment_processed", body)
}
func (b *MsgBroker) Review(body []byte) error {
	return b.publishMessage("review_submitted", body)
}

func (b *MsgBroker) publishMessage(queueName string, body []byte) error {
	err := b.channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		rmq.Publishing{
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
