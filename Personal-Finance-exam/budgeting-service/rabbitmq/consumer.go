package rabbitmq

import (
	genproto "budgeting/genproto/budget"
	genprot "budgeting/genproto/goal"
	genprotos "budgeting/genproto/transaction"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	logger "budgeting/log"
	"budgeting/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MsgBroker struct {
		service          *service.TransactionIService
		service1         *service.GoalService
		service2         *service.BudgetService
		channel          *amqp.Channel
		create           <-chan amqp.Delivery
		update           <-chan amqp.Delivery
		update1          <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
		Db               *mongo.Database
	}
)

func New(service *service.TransactionIService, goalService *service.GoalService, budgetService *service.BudgetService,
	channel *amqp.Channel,
	create <-chan amqp.Delivery,
	update <-chan amqp.Delivery,
	update1 <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int,
	Db *mongo.Database) *MsgBroker {
	return &MsgBroker{
		service:          service,
		service1:         goalService,
		service2:         budgetService,
		channel:          channel,
		create:           create,
		update:           update,
		update1:          update1,
		logger:           logger.InitLogger(),
		wg:               wg,
		numberOfServices: numberOfServices,
		Db:               Db,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.create, "CreateTransaction")
	go m.consumeMessages(consumerCtx, m.update, "UpdateBudget")
	go m.consumeMessages(consumerCtx, m.update1, "UpdateGoal")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	m.logger.Info("Shutting down, waiting for consumers to finish")
	cancel()
	m.wg.Wait()
	m.logger.Info("All consumers have stopped")
}

func (m *MsgBroker) consumeMessages(ctx context.Context, messages <-chan amqp.Delivery, logPrefix string) {
	defer m.wg.Done()
	for {
		select {
		case val := <-messages:
			var err error

			switch logPrefix {
			case "CreateTransaction":
				var req genprotos.CreateTransactionReq

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service.CreateTransaction(ctx, &req)
				if err != nil {
					m.logger.Error("Error while creating booking", "error", err)
					val.Nack(false, false)
					return
				}
				val.Ack(false)

				fmt.Println(req.String())

			case "UpdateBudget":
				var req genproto.UpdateBudgetReq

				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service2.UpdateBudget(ctx, &req)

				fmt.Println(req.String())

			case "UpdateGoal":
				var req genprot.UpdateGoalReq
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					return
				}
				_, err = m.service1.UpdateGoal(ctx, &req)

				fmt.Println(req.String())

			}

			if err != nil {
				m.logger.Error("Failed in %s: %s", logPrefix, err.Error())
				val.Nack(false, false)
				return
			}

			val.Ack(false)

		case <-ctx.Done():
			m.logger.Info("Context done, stopping consumer", "consumer", logPrefix)
			return
		}
	}
}
