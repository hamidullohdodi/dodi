package hendler

import (
	"api-geteway/genproto/booking"
	"api-geteway/service"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ReviewHandler interface{
	CreateReview(c *gin.Context)
	ListReviews(c *gin.Context)
	UpdateReview(c *gin.Context)
	DeleteReview(c *gin.Context)
}

type reviewHendler struct {
	bookingService booking.BookingClient
	logger         *slog.Logger
}

func NewReviewHendler(serviceManager service.ServiceManager, log *slog.Logger) ReviewHandler {
	bookingServiceClient := serviceManager.BookingService()
	if bookingServiceClient == nil {
		log.Error("BookingServiceClient is nil")
		return nil
	}
	return &reviewHendler{
		bookingService: serviceManager.BookingService(),
		logger:         log,
	}
}

// @Summary CreateReview
// @Description CreateReview
// @Tags review
// @ID create-review
// @Accept  json
// @Produce  json
// @Param review body booking.CreateReviewRequest true "CreateReviewRequest"
// @Success 200 {object} booking.CreateReviewResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /review/create-review [post]
func (b *reviewHendler) CreateReview(c *gin.Context) {
	req := &booking.CreateReviewRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	resp, err := b.bookingService.CreateReview(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary ListReviews
// @Description ListReviews
// @Tags review
// @ID list-reviews
// @Accept  json
// @Produce  json
// @Success 200 {object} booking.ListReviewsResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /review/list-review [get]
func (b *reviewHendler) ListReviews(c *gin.Context) {
	req := &booking.ListReviewsRequest{}
	resp, err := b.bookingService.ListReviews(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary UpdateReview
// @Description UpdateReview
// @Tags review
// @ID update-review
// @Accept  json
// @Produce  json
// @Param review body booking.UpdateReviewRequest true "UpdateReviewRequest"
// @Success 200 {object} booking.UpdateReviewResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /review/update-review [put]
func (b *reviewHendler) UpdateReview(c *gin.Context) {
	req := &booking.UpdateReviewRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := b.bookingService.UpdateReview(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary DeleteReview
// @Description DeleteReview
// @Tags review
// @ID delete-review
// @Accept  json
// @Produce  json
// @Param review body booking.DeleteReviewRequest true "DeleteReviewRequest"
// @Success 200 {object} booking.DeleteReviewResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /review/delete-review [delete]
func (b *reviewHendler) DeleteReview(c *gin.Context) {
	req := &booking.DeleteReviewRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := b.bookingService.DeleteReview(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}