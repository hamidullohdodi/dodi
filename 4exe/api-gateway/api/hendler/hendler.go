package hendler

import (
	"api-geteway/service"
	"log/slog"
)

type Handler struct {
	Booking BookingHandler
	User    UserHandler
	Payment PaymentHandler
	Provide ProviderHandler
	Review  ReviewHandler
	Service ServiceHandler
}

func (h *Handler) NewBookingHandler() BookingHandler {
	return h.Booking
}

func (h *Handler) NewAuthHendler() UserHandler {
	return h.User
}

func (h *Handler) NewPaymentHendler() PaymentHandler {
	return h.Payment
}

func (h *Handler) NewProviderHendler() ProviderHandler {
	return h.Provide
}

func (h *Handler) NewReviewHendler() ReviewHandler {
	return h.Review
}

func (h *Handler) NewServiceHendler() ServiceHandler {
	return h.Service
}

type MainHendler interface {
	NewBookingHandler() BookingHandler
	NewAuthHendler() UserHandler
	NewPaymentHendler() PaymentHandler
	NewProviderHendler() ProviderHandler
	NewReviewHendler() ReviewHandler
	NewServiceHendler() ServiceHandler
}

func NewMainHandler(serviceManager service.ServiceManager, log *slog.Logger) MainHendler {
	return &Handler{
		Booking: NewBookingHandler(serviceManager, log),
		User:    NewUserHendler(serviceManager, log),
		Payment: NewPaymentHendler(serviceManager, log),
		Provide: NewProviderHendler(serviceManager, log),
		Review:  NewReviewHendler(serviceManager, log),
		Service: NewServiceHendler(serviceManager, log),
	}
}
