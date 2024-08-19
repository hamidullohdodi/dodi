package notification

import (
  "budgeting/genproto/notification"
  "budgeting/storage"
  "context"
  "fmt"
  "github.com/google/uuid"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "log/slog"
  "time"
)

type AccountRepoI struct {
  Coll *mongo.Collection
  Log  *slog.Logger
}

func NewAccountRepoI(mdb *mongo.Database, log *slog.Logger) storage.Redis {
  return &AccountRepoI{
    Coll: mdb.Collection("Notfk"),
    Log:  log,
  }
}

func (b *AccountRepoI) CreateNotification(req *notification.CreateNotificationReq) (*notification.NotificationResp, error) {
  account := bson.M{
    "user_id":    uuid.New().String(),
    "title":      req.Title,
    "massage":    req.Message,
    "created_at": time.Now().String(),
  }

  _, err := b.Coll.InsertOne(context.Background(), account)
  if err != nil {
    b.Log.Error("Failed to create account", "error", err)
    return nil, err
  }

  return &notification.NotificationResp{
    Notification: &notification.Notification{
      CreatedAt: time.Now().String(),
    },
  }, nil
}
func (b *AccountRepoI) GetNotification(req *notification.GetNotificationReq) (*notification.NotificationResp, error) {
  var account bson.M

  err := b.Coll.FindOne(context.Background(), bson.M{"user_id": req.Id}).Decode(&account)
  if err == mongo.ErrNoDocuments {
    return nil, fmt.Errorf("account not found")
  } else if err != nil {
    b.Log.Error("Failed to get account", "error", err)
    return nil, err
  }

  return &notification.NotificationResp{
    Notification: &notification.Notification{
      Id:        account["_id"].(primitive.ObjectID).Hex(),
      Title:     account["title"].(string),
      Readen:    true,
      CreatedAt: time.Now().String(),
    },
  }, nil
}
func (b *AccountRepoI) UpdateNotification(req *notification.UpdateNotificationReq) (*notification.NotificationResp, error) {

  filter := bson.M{"user_id": req.Id}
  update := bson.M{
    "$set": bson.M{
      "title":       req.Title,
      "description": req.Message,
    },
  }

  result := b.Coll.FindOneAndUpdate(context.Background(), filter, update)
  if result.Err() != nil {
    if result.Err() == mongo.ErrNoDocuments {
      return nil, fmt.Errorf("account not found")
    }
    b.Log.Error("Failed to update account", "error", result.Err())
    return nil, result.Err()
  }

  var updatedAccount notification.NotificationResp
  if err := result.Decode(&updatedAccount); err != nil {
    b.Log.Error("Failed to decode updated account", "error", err)
    return nil, err
  }

  return &updatedAccount, nil
}

func (b *AccountRepoI) DeleteNotification(req *notification.DeleteNotificationReq) (*notification.DeleteNotificationResp, error) {
  result := b.Coll.FindOneAndDelete(context.Background(), bson.M{"user_id": req.Id})
  if result.Err() == mongo.ErrNoDocuments {
    return nil, fmt.Errorf("account not found")
  } else if result.Err() != nil {
    b.Log.Error("Failed to delete account", "error", result.Err())
    return nil, result.Err()
  }

  return &notification.DeleteNotificationResp{}, nil
}
