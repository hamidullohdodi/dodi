package budget

import (
	pb "budgeting/genproto/budget"
	"budgeting/storage"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"strconv"
	"time"
)

type BudgetRepo struct {
	Coll *mongo.Collection
	Log  *slog.Logger
}

func NewBudgetRepo(mdb *mongo.Database, log *slog.Logger) storage.Budget {
	return &BudgetRepo{
		Coll: mdb.Collection("Budget"),
		Log:  log,
	}
}

func (b *BudgetRepo) CreateBudget(req *pb.CreateBudgetReq) (*pb.BudgetResp, error) {
	budget := bson.M{
		"user_id":    req.UserId,
		"name":       req.Name,
		"amount":     req.Amount,
		"start_date": req.StartingTime,
		"end_date":   req.EndingTime,
		"created_at": time.Now().Format(time.RFC3339),
	}

	result, err := b.Coll.InsertOne(context.Background(), budget)
	if err != nil {
		b.Log.Error("Failed to create budget", "error", err)
		return nil, err
	}

	return &pb.BudgetResp{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		UserId:    req.UserId,
		Name:      req.Name,
		Amount:    req.Amount,
		StartingTime: req.StartingTime,
		EndingTime:   req.EndingTime,
		CreatedAt: budget["created_at"].(string),
	}, nil
}
func (b *BudgetRepo) GetBudget(req *pb.GetBudgetReq) (*pb.BudgetResp, error) {

	filter := bson.M{"user_id": bson.M{"$eq": req.UserId}}

	var budget bson.M
	err := b.Coll.FindOne(context.Background(), filter).Decode(&budget)
	if err != nil {
		b.Log.Error("Failed to get budget", "error", err)
		return nil, err
	}

	return &pb.BudgetResp{
		Id:        budget["_id"].(primitive.ObjectID).Hex(),
		UserId:    req.UserId,
		Name:      budget["name"].(string),
		Amount:    budget["amount"].(string),
		StartingTime: budget["starting_time"].(string),
		EndingTime:   budget["ending_time"].(string),
		CreatedAt: budget["created_at"].(string),
	}, nil
}
func (b *BudgetRepo) UpdateBudget(req *pb.UpdateBudgetReq) (*pb.BudgetResp, error) {
	update := bson.M{
		"$set": bson.M{
			"name":       req.Name,
			"amount":     req.Amount,
			"start_date": req.StartingTime,
			"end_date":   req.EndingTime,
		},
	}

	filter := bson.M{"_id": bson.M{"$eq": req.UserId}}
	result, err := b.Coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		b.Log.Error("Failed to update budget", "error", err)
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no budget updated")
	}

	return &pb.BudgetResp{
		Id:        req.Id,
		UserId:    req.UserId,
		Name:      req.Name,
		Amount:    req.Amount,
		StartingTime: req.StartingTime,
		EndingTime:   req.EndingTime,
	}, nil
}
func (b *BudgetRepo) DeleteBudget(req *pb.DeleteBudgetReq) (*pb.Void, error) {
	filter := bson.M{"user_id": bson.M{"$eq": req.UserId}}

	result, err := b.Coll.DeleteOne(context.Background(), filter)
	if err != nil {
		b.Log.Error("Failed to delete budget", "error", err)
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("no budget deleted")
	}

	return &pb.Void{}, nil
}
func (b *BudgetRepo) ListBudgets(req *pb.ListBudgetsReq) (*pb.ListBudgetsResp, error) {
	filter := bson.D{}

	limit, err := strconv.ParseInt(req.Limit, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limit: %v", err)
	}

	paid, err := strconv.ParseInt(req.Paid, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset: %v", err)
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(paid)

	cursor, err := b.Coll.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find budgets: %v", err)
	}
	defer cursor.Close(context.Background())

	var budgets []*pb.BudgetResp
	for cursor.Next(context.Background()) {
		var budget pb.BudgetResp
		if err := cursor.Decode(&budget); err != nil {
			return nil, fmt.Errorf("failed to decode budget: %v", err)
		}
		budgets = append(budgets, &budget)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	res := &pb.ListBudgetsResp{
		Budgets: budgets,
	}

	return res, nil
}
