package account

import (
	pb "budgeting/genproto/account"
	"budgeting/storage"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"strconv"
	"time"
)

type AccountRepo struct {
	Coll *mongo.Collection
	Log  *slog.Logger
}

func NewAccountRepo(mdb *mongo.Database, log *slog.Logger) storage.AccountI {
	return &AccountRepo{
		Coll: mdb.Collection("Account"),
		Log:  log,
	}
}

func (b *AccountRepo) CreateAccount(req *pb.CreateAccountReq) (*pb.AccountResp, error) {
	account := bson.M{
		"user_id":     req.UserId,
		"name":        req.Name,
		"description": req.Description,
		"created_at":  time.Now().String(),
	}

	result, err := b.Coll.InsertOne(context.Background(), account)
	if err != nil {
		b.Log.Error("Failed to create account", "error", err)
		return nil, err
	}

	return &pb.AccountResp{
		Id:          result.InsertedID.(primitive.ObjectID).Hex(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   account["created_at"].(string),
	}, nil
}
func (b *AccountRepo) GetAccount(req *pb.GetAccountReq) (*pb.AccountResp, error) {
	var account bson.M

	err := b.Coll.FindOne(context.Background(), bson.M{"user_id": req.UserId}).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("account not found")
	} else if err != nil {
		b.Log.Error("Failed to get account", "error", err)
		return nil, err
	}

	return &pb.AccountResp{
		Id:          account["_id"].(primitive.ObjectID).Hex(),
		Name:        account["name"].(string),
		Description: account["description"].(string),
		CreatedAt:   account["created_at"].(string),
	}, nil
}
func (b *AccountRepo) UpdateAccount(req *pb.UpdateAccountReq) (*pb.AccountResp, error) {

	filter := bson.M{"user_id": req.UserId}
	update := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
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

	var updatedAccount pb.AccountResp
	if err := result.Decode(&updatedAccount); err != nil {
		b.Log.Error("Failed to decode updated account", "error", err)
		return nil, err
	}

	return &updatedAccount, nil
}

func (b *AccountRepo) DeleteAccount(req *pb.DeleteAccountReq) (*pb.Void1, error) {
	result := b.Coll.FindOneAndDelete(context.Background(), bson.M{"user_id": req.UserId})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("account not found")
	} else if result.Err() != nil {
		b.Log.Error("Failed to delete account", "error", result.Err())
		return nil, result.Err()
	}

	return &pb.Void1{}, nil
}

func (b *AccountRepo) ListAccounts(req *pb.ListAccountsReq) (*pb.ListAccountsResp, error) {
	filter := bson.D{}

	limit, err := strconv.ParseInt(req.Limit, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limit: %v", err)
	}

	offset, err := strconv.ParseInt(req.Paid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset: %v", err)
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	cursor, err := b.Coll.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find accounts: %v", err)
	}
	defer cursor.Close(context.Background())

	var accounts []*pb.AccountResp
	for cursor.Next(context.Background()) {
		var acc pb.AccountResp
		if err := cursor.Decode(&acc); err != nil {
			return nil, fmt.Errorf("failed to decode account: %v", err)
		}
		accounts = append(accounts, &acc)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	res := &pb.ListAccountsResp{
		Accounts: accounts,
	}

	return res, nil
}
