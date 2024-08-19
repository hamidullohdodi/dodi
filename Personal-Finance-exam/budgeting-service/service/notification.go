package service

import (
  "budgeting/genproto/notification"
  "budgeting/storage"
  "context"
  "log/slog"
)

type NotificationService struct {
  notification.UnimplementedNotificationServiceServer
  storage storage.Redis
  logger  *slog.Logger
}

func NewNotificationServiceRepo(storage storage.Redis, logger *slog.Logger) *NotificationService {
  return &NotificationService{
    storage: storage,
    logger:  logger,
  }
}

func (n *NotificationService) CreateNotification(ctx context.Context, req *notification.CreateNotificationReq) (*notification.NotificationResp, error) {
  resp, err := n.storage.CreateNotification(req)
  if err != nil {
    n.logger.Error("Failed to create notification")
    return nil, err
  }
  return resp, nil
}

func (n *NotificationService) UpdateNotification(ctx context.Context, req *notification.UpdateNotificationReq) (*notification.NotificationResp, error) {
  resp, err := n.storage.UpdateNotification(req)
  if err != nil {
    n.logger.Error("Failed to update notification")
    return nil, err
  }
  return resp, nil
}

func (n *NotificationService) GetNotification(ctx context.Context, req *notification.GetNotificationReq) (*notification.NotificationResp, error) {
  resp, err := n.storage.GetNotification(req)
  if err != nil {
    n.logger.Error("Failed to get notification")
    return nil, err
  }
  return resp, nil
}

func (n *NotificationService) DeleteNotification(ctx context.Context, req *notification.DeleteNotificationReq) (*notification.DeleteNotificationResp, error) {
  resp, err := n.storage.DeleteNotification(req)
  if err != nil {
    n.logger.Error("Failed to delete notification")
    return nil, err
  }
  return resp, nil
}
