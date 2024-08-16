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

type ServiceHandler interface {
	CreateService(c *gin.Context)
	UpdateService(c *gin.Context)
	DeleteService(c *gin.Context)
	SearchServices(c *gin.Context)
	ListServices(c *gin.Context)
}

type serviceHendler struct {
	bookingService booking.BookingClient
	logger         *slog.Logger
}

func NewServiceHendler(serviceManager service.ServiceManager, log *slog.Logger) ServiceHandler {
	bookingServiceClient := serviceManager.BookingService()
	if bookingServiceClient == nil {
		log.Error("BookingServiceClient is nil")
		return nil
	}
	return &serviceHendler{
		bookingService: serviceManager.BookingService(),
		logger:         log,
	}
}

// @Summary CreateService
// @Description CreateService
// @Tags service
// @ID create-service
// @Accept  json
// @Produce  json
// @Param service body booking.CreateServiceRequest true "CreateServiceRequest"
// @Success 200 {object} booking.CreateServiceResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /service/create-service [post]
func (h *serviceHendler) CreateService(c *gin.Context) {
	h.logger.Info("CreateService request")
	var req booking.CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind json", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.bookingService.CreateService(context.Background(), &req)
	if err != nil {
		h.logger.Error("Failed to create service", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary UpdateService
// @Description UpdateService
// @Tags service
// @ID update-service
// @Accept  json
// @Produce  json
// @Param service body booking.UpdateServiceRequest true "UpdateServiceRequest"
// @Success 200 {object} booking.UpdateServiceResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /service/update-service [put]
func (h *serviceHendler) UpdateService(c *gin.Context) {
	h.logger.Info("UpdateService request")
	var req booking.UpdateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind json", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.bookingService.UpdateService(context.Background(), &req)
	if err != nil {
		h.logger.Error("Failed to update service", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary DeleteService
// @Description DeleteService
// @Tags service
// @ID delete-service
// @Accept  json
// @Produce  json
// @Param service body booking.DeleteServiceRequest true "DeleteServiceRequest"
// @Success 200 {object} booking.DeleteServiceResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /service/delete-service [delete]
func (h *serviceHendler) DeleteService(c *gin.Context) {
	h.logger.Info("DeleteService request")
	var req booking.DeleteServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind json", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.bookingService.DeleteService(context.Background(), &req)
	if err != nil {
		h.logger.Error("Failed to delete service", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary ListServices
// @Description ListServices
// @Tags service
// @ID list-service
// @Accept  json
// @Produce  json
// @Success 200 {object} booking.ListServicesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /service/list-service [get]
func (h *serviceHendler) ListServices(c *gin.Context) {
	h.logger.Info("ListServices request")
	resp, err := h.bookingService.ListServices(context.Background(), &booking.ListServicesRequest{})
	if err != nil {
		h.logger.Error("Failed to list service", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary SearchServices
// @Description Search for services with optional filters
// @Tags service
// @ID search-service
// @Accept  json
// @Produce  json
// @Param id query string false "Service ID"
// @Param user_id query string false "User ID"
// @Param price query number false "Price"
// @Param duration query int false "Duration"
// @Param description query string false "Description"
// @Success 200 {object} booking.SearchServicesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /service/search-service [get]
func (h *serviceHendler) SearchServices(c *gin.Context) {
	h.logger.Info("SearchServices request")

	req := &booking.SearchServicesRequest{
		Id:          c.Query("id"),
		UserId:      c.Query("user_id"),
		Price:       parseFloatQuery(c, "price"),
		Duration:    parseIntQuery(c, "duration"),
		Description: c.Query("description"),
	}

	resp, err := h.bookingService.SearchServices(context.Background(), req)
	if err != nil {
		h.logger.Error("Failed to search service", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Helper functions to parse query parameters
func parseFloatQuery(c *gin.Context, key string) float32 {
	value := c.Query(key)
	if value == "" {
		return 0
	}
	parsed, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}
	return float32(parsed)
}

func parseIntQuery(c *gin.Context, key string) int32 {
	value := c.Query(key)
	if value == "" {
		return 0
	}
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int32(parsed)
}
