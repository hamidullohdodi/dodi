package handler

import (
	"api/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type Handler struct {
	User         UserHandler
	Account      AccountHandler
	Budger       BudgetHandler
	Goal         GoalHandler
	Transaction  TransactionHandler
	Category     CategoryHandler
	Producer     *service.MsgBroker
	Notification Notification
}

func (h *Handler) NewUserHandler() UserHandler {
	return h.User
}

func (h *Handler) NewNotificationSSS() Notification {
	return h.Notification
}

func (h *Handler) NewAccountHandler() AccountHandler {
	return h.Account
}

func (h *Handler) NewBudgetHandler() BudgetHandler {
	return h.Budger
}

func (h *Handler) NewGoalHandler() GoalHandler {
	return h.Goal
}

func (h *Handler) NewTransactionHandler() TransactionHandler {
	return h.Transaction
}

func (h *Handler) NewCategoryHandler() CategoryHandler {
	return h.Category
}

type MainHandler interface {
	NewUserHandler() UserHandler
	NewAccountHandler() AccountHandler
	NewBudgetHandler() BudgetHandler
	NewGoalHandler() GoalHandler
	NewTransactionHandler() TransactionHandler
	NewCategoryHandler() CategoryHandler
	NewNotificationSSS() Notification
}

func NewMainHandler(serviceManger service.ServiceManager, logger *slog.Logger, conn *amqp.Channel) MainHandler {
	return &Handler{
		User:         NewUserHandler(serviceManger, logger),
		Account:      NewAccountHandler(serviceManger, logger),
		Budger:       NewBudgetHandler(serviceManger, logger, conn),
		Goal:         NewGoalHandler(serviceManger, logger, conn),
		Transaction:  NewTransactionHandler(serviceManger, logger, conn),
		Category:     NewCategoryHandler(serviceManger, logger),
		Producer:     service.NewMsgBroker(conn, logger),
		Notification: NewNotificationSSS(serviceManger, logger),
	}
}
