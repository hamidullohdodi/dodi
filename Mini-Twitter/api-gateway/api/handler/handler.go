package handler

import (
	"apigateway/service"
	"log/slog"
)

type Handler struct {
	User UserHandler
}

func (h *Handler) NewHandler() UserHandler {
	return h.User
}

type handler interface {
	NewHandler() UserHandler
}

func NewMainHandler(service service.Service, logger *slog.Logger) handler {
	return &Handler{
		User: NewUserHandler(service, logger),
	}
}
