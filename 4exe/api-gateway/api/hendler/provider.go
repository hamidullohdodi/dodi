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

type ProviderHandler interface {
	CreateProvider(c *gin.Context)
	ListProvider(c *gin.Context)
	UpdateProvider(c *gin.Context)
	DeleteProvider(c *gin.Context)
	GetProvider(c *gin.Context)
	SearchProviders(c *gin.Context)
}

type providerHendler struct {
	bookingService booking.BookingClient
	logger         *slog.Logger
}

func NewProviderHendler(serviceManager service.ServiceManager, logger *slog.Logger) ProviderHandler {
	bookingServiceClient := serviceManager.BookingService()
	if bookingServiceClient == nil {
		logger.Error("BookingServiceClient is nil")
		return nil
	}
	return &providerHendler{
		bookingService: bookingServiceClient,
		logger:         logger,
	}
}

// @Summary CreateProvider
// @Description Create a new provider
// @Tags provider
// @ID create-provider
// @Accept json
// @Produce json
// @Param provider body booking.CreateProviderRequest true "CreateProviderRequest"
// @Success 200 {object} booking.CreateProviderResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/create-provider [post]
func (h *providerHendler) CreateProvider(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &booking.CreateProviderRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.bookingService.CreateProvider(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary ListProviders
// @Description List providers with pagination
// @Tags provider
// @ID list-providers
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} booking.ListProvidersResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/list-providers [get]
func (h *providerHendler) ListProvider(c *gin.Context) {
	req := &booking.ListProvidersRequest{}
	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Limit = int32(limit)
	}
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Offset = int32(offset)
	}

	resp, err := h.bookingService.ListProviders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary UpdateProvider
// @Description Update an existing provider
// @Tags provider
// @ID update-provider
// @Accept json
// @Produce json
// @Param provider body booking.UpdateProviderRequest true "UpdateProviderRequest"
// @Success 200 {object} booking.UpdateProviderResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/update-provider [put]
func (h *providerHendler) UpdateProvider(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &booking.UpdateProviderRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.bookingService.UpdateProvider(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary DeleteProvider
// @Description Delete a provider by ID
// @Tags provider
// @ID delete-provider
// @Accept json
// @Produce json
// @Param id query string true "Provider ID"
// @Success 200 {object} booking.DeleteProviderResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/delete-provider [delete]
func (h *providerHendler) DeleteProvider(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	resp, err := h.bookingService.DeleteProvider(ctx, &booking.DeleteProviderRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary GetProvider
// @Description Retrieve a provider by ID
// @Tags provider
// @ID get-provider
// @Accept json
// @Produce json
// @Param id path string true "Provider ID"
// @Success 200 {object} booking.GetProviderResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/{id} [get]
func (h *providerHendler) GetProvider(c *gin.Context) {
	id := c.Param("id")
	req := &booking.GetProviderRequest{Id: id}
	resp, err := h.bookingService.GetProvider(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary SearchProviders
// @Description Search providers based on various criteria such as user ID, company name, availability, and location.
// @Tags provider
// @ID search-providers
// @Accept json
// @Produce json
// @Param userId query string false "User ID" example("12345")
// @Param companyName query string false "Company Name" example("Tech Solutions")
// @Param city query string false "City" example("San Francisco")
// @Param country query string false "Country" example("USA")
// @Success 200 {object} booking.SearchProvidersResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /provider/search-provider [get]
func (h *providerHendler) SearchProviders(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req := &booking.SearchProvidersRequest{
		UserId:      c.Query("userId"),
		CompanyName: c.Query("companyName"),
		Location: &booking.Location{
			City:    c.Query("city"),
			Country: c.Query("country"),
		},
	}

	resp, err := h.bookingService.SearchProviders(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
