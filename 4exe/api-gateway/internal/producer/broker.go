package producer

import (
	"log"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
  MsgBroker struct {
    channel *amqp.Channel
    logger  *slog.Logger
  }
)

func NewMsgBroker(channel *amqp.Channel, logger *slog.Logger) *MsgBroker {
  return &MsgBroker{
    channel: channel,
    logger:  logger,
  }
}

func (b *MsgBroker) CreateBooking(body []byte) error {
  log.Println("=/=/=/=/=/=/=/=/=/")
  return b.publishMessage("create_booking", body)
}

func (b *MsgBroker) UpdateBooking(body []byte) error {
  log.Println("=/=/=/=/=/=/=/=/=/")
  return b.publishMessage("booking_updated", body)
}
func (b *MsgBroker) CancelBooking(body []byte) error {
  log.Println("=/=/=/=/=/=/=/=/=/")
  return b.publishMessage("booking_cancelled", body)
}
func (b *MsgBroker) Payment(body []byte) error {
  log.Println("=/=/=/=/=/=/=/=/=/")
  return b.publishMessage("payment_processed", body)
}
func (b *MsgBroker) Review(body []byte) error {
  log.Println("=/=/=/=/=/=/=/=/")
  return b.publishMessage("review_submitted", body)
}


func (b *MsgBroker) publishMessage(queueName string, body []byte) error {
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
