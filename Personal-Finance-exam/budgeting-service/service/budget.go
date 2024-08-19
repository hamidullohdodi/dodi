package service

import (
	pb "budgeting/genproto/budget"
	"budgeting/storage"
	"context"
	"fmt"
	"log/slog"
)

type BudgetService struct {
	pb.UnimplementedBudgetServiceServer
	storage.Budget
	logger *slog.Logger
}

func NewBudgetService(storage storage.Budget, logger *slog.Logger) *BudgetService {
	return &BudgetService{
		Budget: storage,
		logger: logger,
	}
}

func (b *BudgetService) CreateBudget(ctx context.Context, req *pb.CreateBudgetReq) (*pb.BudgetResp, error) {
	resp, err := b.Budget.CreateBudget(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not create budget: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *BudgetService) GetBudget(ctx context.Context, req *pb.GetBudgetReq) (*pb.BudgetResp, error) {
	resp, err := b.Budget.GetBudget(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not get budget: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *BudgetService) UpdateBudget(ctx context.Context, req *pb.UpdateBudgetReq) (*pb.BudgetResp, error) {
	resp, err := b.Budget.UpdateBudget(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not update budget: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *BudgetService) DeleteBudget(ctx context.Context, req *pb.DeleteBudgetReq) (*pb.Void, error) {
	_, err := b.Budget.DeleteBudget(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not delete budget: %v", err))
		return nil, err
	}
	return &pb.Void{}, nil
}
func (b *BudgetService) ListBudgets(ctx context.Context, req *pb.ListBudgetsReq) (*pb.ListBudgetsResp, error) {
	resp, err := b.Budget.ListBudgets(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not list budgets: %v", err))
		return nil, err
	}
	return resp, nil
}
