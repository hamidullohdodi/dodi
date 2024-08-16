package hendler

import (
	"api-geteway/genproto/booking"
	"api-geteway/service"
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	CreatePayment(c *gin.Context)
	UpdatePayment(c *gin.Context)
	DeletePayment(c *gin.Context)
	ListPayments(c *gin.Context)
	GetPayment(c *gin.Context)
}

type paymentHendler struct {
	bookingService booking.BookingClient
	logger         *slog.Logger
}

func NewPaymentHendler(serviceManager service.ServiceManager, log *slog.Logger) PaymentHandler {
	bookingServiceClient := serviceManager.BookingService()
	if bookingServiceClient == nil {
		log.Error("BookingServiceClient is nil")
		return nil
	}
	return &paymentHendler{
		bookingService: serviceManager.BookingService(),
		logger:         log,
	}
}

// @Summary CreatePayment
// @Description CreatePayment
// @Tags payment
// @ID create-payment
// @Accept  json
// @Produce  json
// @Param payment body booking.CreatePaymentRequest true "CreatePaymentRequest"
// @Success 200 {object} booking.CreatePaymentResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/create-payment [post]
func (h *paymentHendler) CreatePayment(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &booking.CreatePaymentRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.bookingService.CreatePayment(ctx, req)
	if err != nil {
		h.logger.Error("Failed to create payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary UpdatePayment
// @Description UpdatePayment
// @Tags payment
// @ID update-payment
// @Accept  json
// @Produce  json
// @Param payment body booking.UpdatePaymentRequest true "UpdatePaymentRequest"
// @Success 200 {object} booking.UpdatePaymentResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/update-payment [put]
func (h *paymentHendler) UpdatePayment(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &booking.UpdatePaymentRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.logger.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.bookingService.UpdatePayment(ctx, req)
	if err != nil {
		h.logger.Error("Failed to update payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary DeletePayment
// @Description DeletePayment
// @Tags payment
// @ID delete-payment
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} booking.DeletePaymentResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/delete-payment/{id} [delete]
func (h *paymentHendler) DeletePayment(c *gin.Context) {
	payment_id := c.Param("id")
	// log.Printf("request id: %s", Id)
	res, err := h.bookingService.DeletePayment(c, &booking.DeletePaymentRequest{Id: payment_id})
	if err != nil {
		h.logger.Error("Failed to delete payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @sSummary GetPayment
// @Description GetPayment
// @Tags payment
// @ID get-payment
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} booking.GetPaymentResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/{id} [get]
func (h *paymentHendler) GetPayment(c *gin.Context) {
	Id := c.Query("id")
	if Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	res, err := h.bookingService.GetPayment(c, &booking.GetPaymentRequest{Id: Id})
	if err != nil {
		h.logger.Error("Failed to get payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary ListPayment
// @Description ListPayment
// @Tags payment
// @ID list-payment
// @Accept  json
// @Produce  json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} booking.ListPaymentsResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/list-payments [get]
func (h *paymentHendler) ListPayments(c *gin.Context) {
	req := &booking.ListPaymentsRequest{}
	LimitSrt := c.Query("limit")
	if LimitSrt != "" {
		limit, err := strconv.Atoi(LimitSrt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Limit = int32(limit)
	}
	OffsetSrt := c.Query("offset")
	if OffsetSrt != "" {
		offset, err := strconv.Atoi(OffsetSrt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Offset = int32(offset)
	}
	res, err := h.bookingService.ListPayments(c, req)
	if err != nil {
		h.logger.Error("Failed to list payment", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
