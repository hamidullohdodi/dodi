package transaction

import (
	pb "budgeting/genproto/transaction"
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

type TransactionRepo struct {
	Coll *mongo.Collection
	Log  *slog.Logger
}

func NewTransactionRepo(mdb *mongo.Database, log *slog.Logger) storage.TransactionI {
	return &TransactionRepo{
		Coll: mdb.Collection("Transaction"),
		Log:  log,
	}
}

func (b *TransactionRepo) CreateTransaction(req *pb.CreateTransactionReq) (*pb.TransactionResp, error) {

	transaction := bson.M{
		"user_id":     req.UserId,
		"account_id":  req.AccountId,
		"category_id": req.CategoryId,
		"amount":      req.Amount,
		"date":        req.Date,
		"description": req.Description,
		"created_at":  time.Now().Format(time.RFC3339),
	}

	result, err := b.Coll.InsertOne(context.Background(), transaction)
	if err != nil {
		b.Log.Error("Failed to create transaction", "error", err)
		return nil, err
	}

	return &pb.TransactionResp{
		Id:          result.InsertedID.(primitive.ObjectID).Hex(),
		UserId:      req.UserId,
		AccountId:   req.AccountId,
		CategoryId:  req.CategoryId,
		Amount:      req.Amount,
		Date:        req.Date,
		Description: req.Description,
		CreatedAt:   transaction["created_at"].(string),
	}, nil
}
func (b *TransactionRepo) GetTransaction(req *pb.GetTransactionReq) (*pb.TransactionResp, error) {
	var transaction bson.M
	err := b.Coll.FindOne(context.Background(), bson.M{"user_id": req.UserId}).Decode(&transaction)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("transaction not found")
	} else if err != nil {
		b.Log.Error("Failed to get transaction", "error", err)
		return nil, err
	}

	return &pb.TransactionResp{
		Id:          transaction["_id"].(primitive.ObjectID).Hex(),
		UserId:      req.UserId,
		AccountId:   transaction["account_id"].(string),
		CategoryId:  transaction["category_id"].(string),
		Amount:      transaction["amount"].(string),
		Date:        transaction["date"].(string),
		Description: transaction["description"].(string),
		CreatedAt:   transaction["created_at"].(string),
	}, nil
}
func (b *TransactionRepo) UpdateTransaction(req *pb.UpdateTransactionReq) (*pb.TransactionResp, error) {
	update := bson.M{
		"$set": bson.M{
			"account_id":  req.AccountId,
			"category_id": req.CategoryId,
			"amount":      req.Amount,
			"date":        req.Date,
			"description": req.Description,
			"updated_at":  time.Now().Format(time.RFC3339),
		},
	}

	result := b.Coll.FindOneAndUpdate(context.Background(), bson.M{"user_id": req.UserId}, update)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("transaction not found")
	} else if result.Err() != nil {
		b.Log.Error("Failed to update transaction", "error", result.Err())
		return nil, result.Err()
	}

	return &pb.TransactionResp{
		Id:          req.Id,
		UserId:      req.UserId,
		AccountId:   req.AccountId,
		CategoryId:  req.CategoryId,
		Amount:      req.Amount,
		Date:        req.Date,
		Description: req.Description,
		CreatedAt:   time.Now().String(),
	}, nil
}
func (b *TransactionRepo) DeleteTransaction(req *pb.DeleteTransactionReq) (*pb.Void4, error) {
	result := b.Coll.FindOneAndDelete(context.Background(), bson.M{"user_id": req.UserId})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("transaction not found")
	} else if result.Err() != nil {
		b.Log.Error("Failed to delete transaction", "error", result.Err())
		return nil, result.Err()
	}

	return &pb.Void4{}, nil
}

func (b *TransactionRepo) ListTransactions(req *pb.ListTransactionsReq) (*pb.ListTransactionsResp, error) {
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
		return nil, fmt.Errorf("failed to find transactions: %v", err)
	}
	defer cursor.Close(context.Background())

	var transactions []*pb.TransactionResp
	for cursor.Next(context.Background()) {
		var txn pb.TransactionResp
		if err := cursor.Decode(&txn); err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
		}
		transactions = append(transactions, &txn)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	res := &pb.ListTransactionsResp{
		Transactions: transactions,
	}

	return res, nil
}
