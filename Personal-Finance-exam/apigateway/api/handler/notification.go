package handler

import (
  "api/genproto/notification"
  "api/service"
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  "log/slog"
  "net/http"
)

type Notification interface {
  CreateNotification(c *gin.Context)
  GetNotification(c *gin.Context)
  UpdateNotification(c *gin.Context)
  DeleteNotification(c *gin.Context)
}

type NotificationSSS struct {
  NotificationCLent notification.NotificationServiceClient
  logger            *slog.Logger
}

func NewNotificationSSS(serviceManger service.ServiceManager, logger *slog.Logger) Notification {
  return &NotificationSSS{
    NotificationCLent: serviceManger.NotificationService(),
    logger:            logger,
  }
}

// CreateNotification godoc
// @Summary Create a new Notification
// @Description Create a new Notification with the provided details
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Param account body notification.CreateNotificationReq true "Notification Creation Data"
// @Success 201 {object} notification.NotificationResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating Notification "
// @Router /not/create [post]
func (s *NotificationSSS) CreateNotification(c *gin.Context) {
  req := &notification.CreateNotificationReq{}
  if err := c.ShouldBindJSON(req); err != nil {
    s.logger.Error("failed to bind request body")
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  resq, err := s.NotificationCLent.CreateNotification(context.Background(), req)
  if err != nil {
    s.logger.Error("failed to create notification")
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, resq)
}

// GetNotification godoc
// @Summary Get a new Notification
// @Description Get a new Notification with the provided details
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 201 {object} notification.NotificationResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while get Notification "
// @Router /not/get/{id} [get]
func (s *NotificationSSS) GetNotification(c *gin.Context) {
  id := c.Param("id")
  rep := &notification.GetNotificationReq{
    Id: id,
  }
  fmt.Println(id)
  resq, err := s.NotificationCLent.GetNotification(context.Background(), rep)
  if err != nil {
    s.logger.Error("failed to get notification")
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  fmt.Println(id, resq)
  c.JSON(http.StatusOK, resq)
}

// UpdateNotification godoc
// @Summary Update a new Notification
// @Description Update a new Notification with the provided details
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Param account body notification.UpdateNotificationReq true "Notification Creation Data"
// @Success 201 {object} notification.NotificationResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating Notification "
// @Router /not/update [put]
func (s *NotificationSSS) UpdateNotification(c *gin.Context) {
  req := &notification.UpdateNotificationReq{}
  if err := c.ShouldBindJSON(req); err != nil {
    s.logger.Error("failed to bind request body")
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  resq, err := s.NotificationCLent.UpdateNotification(context.Background(), req)
  if err != nil {
    s.logger.Error("failed to update notification")
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, resq)
}


// DeleteNotification godoc
// @Summary Delete a new Notification
// @Description Delete a new Notification with the provided details
// @Security BearerAuth
// @Tags Notification
// @Accept json
// @Produce json
// @Param id query string true "Account ID"
// @Success 201 {object} notification.DeleteNotificationResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Error while creating Notification "
// @Router /not/delete/{id} [delete]
func (s *NotificationSSS) DeleteNotification(c *gin.Context) {
	id := c.Query("id")
  
	rep := notification.DeleteNotificationReq{
	  Id: id,
	}
  
	resq, err := s.NotificationCLent.DeleteNotification(context.Background(), &rep)
	if err != nil {
	  s.logger.Error("failed to delete notification")
	  c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	  return
	}
	c.JSON(http.StatusOK, resq)
  }
  