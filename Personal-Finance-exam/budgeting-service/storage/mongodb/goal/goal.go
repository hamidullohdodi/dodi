package goal

import (
	pb "budgeting/genproto/goal"
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

type GoalRepo struct {
	Coll *mongo.Collection
	Log  *slog.Logger
}

func NewGoalRepo(mdb *mongo.Database, log *slog.Logger) storage.Goal {
	return &GoalRepo{
		Coll: mdb.Collection("Goals"),
		Log:  log,
	}
}

func (b *GoalRepo) CreateGoal(req *pb.CreateGoalReq) (*pb.GoalResp, error) {
	goal := bson.M{
		"name":          req.Name,
		"target_amount": req.TargetAmount,
		"deadline":      req.Deadline,
		"created_at":    time.Now().Format(time.RFC3339),
	}

	result, err := b.Coll.InsertOne(context.Background(), goal)
	if err != nil {
		b.Log.Error("Failed to create goal", "error", err)
		return nil, err
	}

	return &pb.GoalResp{
		UserId:           result.InsertedID.(primitive.ObjectID).Hex(),
		Name:         req.Name,
		TargetAmount: req.TargetAmount,
		Deadline:     req.Deadline,
		CreatedAt:    goal["created_at"].(string),
	}, nil
}

func (b *GoalRepo) GetGoal(req *pb.GetGoalReq) (*pb.GoalResp, error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}
	filter := bson.M{"_id": id}

	var goal bson.M
	err = b.Coll.FindOne(context.Background(), filter).Decode(&goal)
	if err != nil {
		b.Log.Error("Failed to get goal", "error", err)
		return nil, err
	}

	return &pb.GoalResp{
		UserId:           req.UserId,
		Name:         goal["name"].(string),
		TargetAmount: goal["target_amount"].(string),
		Deadline:     goal["deadline"].(string),
		CreatedAt:    goal["created_at"].(string),
	}, nil
}

func (b *GoalRepo) UpdateGoal(req *pb.UpdateGoalReq) (*pb.GoalResp, error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}
	update := bson.M{
		"$set": bson.M{
			"name":           req.Name,
			"target_amount":  req.TargetAmount,
			"current_amount": req.CurrentAmount,
			"deadline":       req.Deadline,
		},
	}

	filter := bson.M{"_id": id}
	result, err := b.Coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		b.Log.Error("Failed to update goal", "error", err)
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("no goal updated")
	}

	return &pb.GoalResp{
		UserId:            req.UserId,
		Name:          req.Name,
		TargetAmount:  req.TargetAmount,
		CurrentAmount: req.CurrentAmount,
		Deadline:      req.Deadline,
	}, nil
}

func (b *GoalRepo) DeleteGoal(req *pb.DeleteGoalReq) (*pb.Void3, error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}
	filter := bson.M{"_id": id}

	result, err := b.Coll.DeleteOne(context.Background(), filter)
	if err != nil {
		b.Log.Error("Failed to delete goal", "error", err)
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("no goal deleted")
	}

	return &pb.Void3{}, nil
}

func (b *GoalRepo) ListGoals(req *pb.ListGoalsReq) (*pb.ListGoalsResp, error) {
	filter := bson.D{}

	limit, err := strconv.ParseInt(req.Limit, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limit: %v", err)
	}

	offset, err := strconv.ParseInt(req.Offset, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset: %v", err)
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	cursor, err := b.Coll.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find goals: %v", err)
	}
	defer cursor.Close(context.Background())

	var goals []*pb.GoalResp
	for cursor.Next(context.Background()) {
		var goal bson.M
		if err := cursor.Decode(&goal); err != nil {
			return nil, fmt.Errorf("failed to decode goal: %v", err)
		}
		goals = append(goals, &pb.GoalResp{
			UserId:           goal["_id"].(primitive.ObjectID).Hex(),
			Name:         goal["name"].(string),
			TargetAmount: goal["target_amount"].(string),
			Deadline:     goal["deadline"].(string),
			CreatedAt:    goal["created_at"].(string),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	res := &pb.ListGoalsResp{
		Goals: goals,
	}

	return res, nil
}

func (b *GoalRepo) GetUserSpending(req *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	collection := b.Coll.Database().Collection("Transaction")

	filter := bson.M{
		"user_id":   req.UserId,
		"type":      "spending",
		"timestamp": bson.M{"$gte": req.StartingTime, "$lte": req.EndingTime},
	}

	result := &pb.GetUserMoneyResponse{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		b.Log.Error("Failed to get user spending", "error", err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var item pb.GetUserMoneyResponse
		err := cur.Decode(&item)
		if err != nil {
			b.Log.Error("Failed to decode user spending", "error", err)
			return nil, err
		}
		result.TotalAmount += item.TotalAmount
		result.CategoryId = item.CategoryId
		result.Time = item.Time
	}

	if err := cur.Err(); err != nil {
		b.Log.Error("Error iterating over cursor", "error", err)
		return nil, err
	}

	return result, nil
}

func (b *GoalRepo) GetUserIncome(req *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	collection := b.Coll.Database().Collection("Transaction")

	filter := bson.M{
		"user_id":   req.UserId,
		"type":      "income",
		"timestamp": bson.M{"$gte": req.StartingTime, "$lte": req.EndingTime},
	}

	result := &pb.GetUserMoneyResponse{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		b.Log.Error("Failed to get user income", "error", err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var item pb.GetUserMoneyResponse
		err := cur.Decode(&item)
		if err != nil {
			b.Log.Error("Failed to decode user income", "error", err)
			return nil, err
		}
		result.TotalAmount += item.TotalAmount
		result.CategoryId = item.CategoryId
		result.Time = item.Time
	}

	if err := cur.Err(); err != nil {
		b.Log.Error("Error iterating over cursor", "error", err)
		return nil, err
	}

	return result, nil
}

func (b *GoalRepo) GetGoalReportProgress(req *pb.GoalProgressRequest) (*pb.GoalProgressResponse, error) {
	collection := b.Coll.Database().Collection("goals")

	filter := bson.M{
		"user_id": req.UserId,
		"timestamp": bson.M{
			"$gte": req.StartingTime,
			"$lte": req.EndingTime,
		},
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		b.Log.Error("Failed to get goal report progress", "error", err)
		return nil, err
	}
	defer cur.Close(context.TODO())

	response := &pb.GoalProgressResponse{}
	for cur.Next(context.TODO()) {
		var item pb.GoalProgressItem
		err := cur.Decode(&item)
		if err != nil {
			b.Log.Error("Failed to decode goal progress item", "error", err)
			return nil, err
		}
		response.Results = append(response.Results, &item)
	}

	if err := cur.Err(); err != nil {
		b.Log.Error("Error iterating over cursor", "error", err)
		return nil, err
	}

	return response, nil
}

func (b *GoalRepo) GetBudgetSummary(req *pb.UserId) (*pb.GetUserBudgetResponse, error) {
	collection := b.Coll.Database().Collection("user")

	filter := bson.M{"user_id": req.UserId}

	var summary pb.GetUserBudgetResponse
	err := collection.FindOne(context.Background(), filter).Decode(&summary)
	if err != nil {
		b.Log.Error("Failed to get budget summary", "error", err)
		return nil, err
	}

	return &summary, nil
}
