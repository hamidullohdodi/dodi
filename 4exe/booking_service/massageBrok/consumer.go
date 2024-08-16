package massagebrok

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	genprotos "booking_service/genproto/booking"
	gg "booking_service/genproto/user"
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
		BookingService   genprotos.BookingClient
		AuthService      gg.AuthClient
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
		service: service,
		channel: channel,
		create:  create,
		delete:  delete,

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

func (m *MsgBroker) consumeMessages(ctx context.Context, deliveries <-chan amqp.Delivery, queueName string) {
	defer m.wg.Done()

	for {
		select {
		case val := <-deliveries:
			var err error
			switch queueName {
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
				if err != nil {
					m.logger.Error("Error while canceling booking", "error", err)
					val.Nack(false, false)
					continue
				}
				val.Ack(false)
			}
		}
	}
}

func (m *MsgBroker) CreateBooking(ctx context.Context, req *genprotos.CreateBookingRequest) (*genprotos.CreateBookingResponse, error) {

	_, err := m.BookingService.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return &genprotos.CreateBookingResponse{}, nil
}

func (m *MsgBroker) CancelBooking(ctx context.Context, req *genprotos.CancelBookingRequest) (*genprotos.CancelBookingResponse, error) {

	_, err := m.BookingService.CancelBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return &genprotos.CancelBookingResponse{}, nil
}
