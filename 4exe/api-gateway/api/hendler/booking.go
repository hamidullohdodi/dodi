package hendler

import (
	"api-geteway/genproto/booking"

	"api-geteway/service"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler interface {
	CreateBooking(c *gin.Context)
	GetBooking(c *gin.Context)
	UpdateBooking(c *gin.Context)
	DeleteBooking(c *gin.Context)
	ListBookings(c *gin.Context)
}

type bookingHendler struct {
	bookingService booking.BookingClient
	logger         *slog.Logger
}

func NewBookingHandler(serviceManager service.ServiceManager, log *slog.Logger) BookingHandler {
	bookingServiceClient := serviceManager.BookingService()
	if bookingServiceClient == nil {
		log.Error("BookingServiceClient is nil")
		return nil
	}
	return &bookingHendler{
		bookingService: serviceManager.BookingService(),
		logger:         log,
	}
}

// @Summary CreateBooking
// @Description CreateBooking
// @Tags booking
// @ID create-booking
// @Accept  json
// @Produce  json
// @Param booking body booking.CreateBookingRequest true "CreateBookingRequest"
// @Success 200 {object} booking.CreateBookingResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /booking/create-booking [post]
func (b *bookingHendler) CreateBooking(c *gin.Context) {
	req := &booking.CreateBookingRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := b.bookingService.CreateBooking(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary GetBooking
// @Description GetBooking
// @Tags booking
// @ID get-booking
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} booking.GetBookingResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /booking/{id} [get]
func (b *bookingHendler) GetBooking(c *gin.Context) {
	req := &booking.GetBookingRequest{
		Id: c.Param("id"),
	}
	resp, err := b.bookingService.GetBooking(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary UpdateBooking
// @Description UpdateBooking
// @Tags booking
// @ID update-booking
// @Accept  json
// @Produce  json
// @Param booking body booking.UpdateBookingRequest true "UpdateBookingRequest"
// @Success 200 {object} booking.UpdateBookingResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /booking/update-booking [put]
func (b *bookingHendler) UpdateBooking(c *gin.Context) {
	req := &booking.UpdateBookingRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := b.bookingService.UpdateBooking(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary DeleteBooking
// @Description DeleteBooking
// @Tags booking
// @ID delete-booking
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} booking.CancelBookingResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /booking/{id} [delete]
func (b *bookingHendler) DeleteBooking(c *gin.Context) {
	req := &booking.CancelBookingRequest{
		Id: c.Param("id"),
	}
	resp, err := b.bookingService.CancelBooking(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary ListBookings
// @Description ListBookings
// @Tags booking
// @ID list-bookings
// @Accept  json
// @Produce  json
// @Success 200 {object} booking.ListBookingsResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /booking/list-bookings [get]
func (b *bookingHendler) ListBookings(c *gin.Context) {
	req := &booking.ListBookingsRequest{}
	resp, err := b.bookingService.ListBookings(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
