package category

import (
	pb "budgeting/genproto/category"
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

type CategoryRepo struct {
	Coll *mongo.Collection
	Log  *slog.Logger
}

func NewCategoryRepo(mdb *mongo.Database, log *slog.Logger) storage.CategoryI {
	return &CategoryRepo{
		Coll: mdb.Collection("Category"),
		Log:  log,
	}
}

func (b *CategoryRepo) CreateCategory(req *pb.CreateCategoryReq) (*pb.CategoryResp, error) {
	category := bson.M{
		"name":        req.Name,
		"description": req.Description,
		"created_at":  time.Now().String(),
	}

	result, err := b.Coll.InsertOne(context.Background(), category)
	if err != nil {
		b.Log.Error("Failed to create category", "error", err)
		return nil, err
	}

	return &pb.CategoryResp{
		Id:          result.InsertedID.(primitive.ObjectID).Hex(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   category["created_at"].(string),
	}, nil
}
func (b *CategoryRepo) UpdateCategory(req *pb.UpdateCategoryReq) (*pb.CategoryResp, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"description": req.Description,
		},
	}

	result := b.Coll.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("category not found")
		}
		b.Log.Error("Failed to update category", "error", result.Err())
		return nil, result.Err()
	}

	var updatedCategory bson.M
	if err := result.Decode(&updatedCategory); err != nil {
		b.Log.Error("Failed to decode updated category", "error", err)
		return nil, err
	}

	return &pb.CategoryResp{
		Id:          updatedCategory["_id"].(primitive.ObjectID).Hex(),
		Name:        updatedCategory["name"].(string),
		Description: updatedCategory["description"].(string),
	}, nil
}

func (b *CategoryRepo) DeleteCategory(req *pb.DeleteCategoryReq) (*pb.Void2, error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid account ID: %v", err)
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}

	result, err := b.Coll.DeleteOne(context.Background(), filter)
	if err != nil {
		b.Log.Error("Failed to delete category", "error", err)
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("no category deleted")
	}

	return &pb.Void2{}, nil
}

func (b *CategoryRepo) ListCategories(req *pb.ListCategoriesReq) (*pb.ListCategoriesResp, error) {
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
		return nil, fmt.Errorf("failed to find categories: %v", err)
	}
	defer cursor.Close(context.Background())

	var categories []*pb.CategoryResp
	for cursor.Next(context.Background()) {
		var cat pb.CategoryResp
		if err := cursor.Decode(&cat); err != nil {
			return nil, fmt.Errorf("failed to decode category: %v", err)
		}
		categories = append(categories, &cat)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	res := &pb.ListCategoriesResp{
		Categories: categories,
	}

	return res, nil
}
