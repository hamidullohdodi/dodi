package service

import (
	pb "budgeting/genproto/transaction"
	"budgeting/storage"
	"context"
	"fmt"
	"log/slog"
)

type TransactionIService struct {
	pb.UnimplementedTransactionServiceServer
	storage.TransactionI
	logger *slog.Logger
}

func NewTransactionIService(storage storage.TransactionI, logger *slog.Logger) *TransactionIService {
	return &TransactionIService{
		TransactionI: storage,
		logger:       logger,
	}
}

func (b *TransactionIService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionReq) (*pb.TransactionResp, error) {
	resp, err := b.TransactionI.CreateTransaction(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not create transaction: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *TransactionIService) GetTransaction(ctx context.Context, req *pb.GetTransactionReq) (*pb.TransactionResp, error) {
	resp, err := b.TransactionI.GetTransaction(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not get transaction: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *TransactionIService) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionReq) (*pb.TransactionResp, error) {
	resp, err := b.TransactionI.UpdateTransaction(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not update transaction: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *TransactionIService) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionReq) (*pb.Void4, error) {
	_, err := b.TransactionI.DeleteTransaction(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not delete transaction: %v", err))
		return nil, err
	}
	return &pb.Void4{}, nil
}
func (b *TransactionIService) ListTransactions(ctx context.Context, req *pb.ListTransactionsReq) (*pb.ListTransactionsResp, error) {
	resp, err := b.TransactionI.ListTransactions(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not list transactions: %v", err))
		return nil, err
	}
	return resp, nil
}
