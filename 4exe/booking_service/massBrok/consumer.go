package messbroker

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"booking_service/cl-logger/logger"
	genprotos "booking_service/genproto/booking"
	"booking_service/service"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MsgBroker struct {
		service          *service.BookingService
		channel          *amqp.Channel
		create           <-chan amqp.Delivery
		delete           <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
		Db               *mongo.Database
	}
)

func New(service *service.BookingService,
	channel *amqp.Channel,
	create <-chan amqp.Delivery,
	delete <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int,
	Db *mongo.Database) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		create:           create,
		delete:           delete,
		logger:           logger.NewLogger(),
		wg:               wg,
		numberOfServices: numberOfServices,
		Db:               Db,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.create, "create_booking")
	go m.consumeMessages(consumerCtx, m.delete, "booking_cancelled")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	cancel()
	m.wg.Wait()
}

func (m *MsgBroker) consumeMessages(ctx context.Context, messages <-chan amqp.Delivery, logPrefix string) {
	defer m.wg.Done()
	for {
		select {
		case val := <-messages:
			var err error

			switch logPrefix {
			case "create_booking":
				var req genprotos.CreateBookingRequest

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}

				_, err = m.service.CreateBooking(ctx, &req)
				if err != nil {
					m.logger.Error("Error while creating booking", "error", err)
					val.Nack(false, false)
					continue
				}
				val.Ack(false)
			case "booking_cancelled":
				var req genprotos.CancelBookingRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				_, err = m.service.CancelBooking(ctx, &req)

			}

			if err != nil {
				m.logger.Error("Failed in %s: %s", logPrefix, err.Error())
				val.Nack(false, false)
				continue
			}

			val.Ack(false)

		case <-ctx.Done():
			m.logger.Info("Context done, stopping consumer", "consumer", logPrefix)
			return
		}
	}
}
