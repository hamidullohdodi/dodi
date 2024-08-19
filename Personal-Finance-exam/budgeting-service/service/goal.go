package service

import (
	pb "budgeting/genproto/goal"
	"budgeting/storage"
	"context"
	"fmt"
	"log/slog"
)

type GoalService struct {
	pb.UnimplementedGoalServiceServer
	storage.Goal
	logger *slog.Logger
}

func NewGoalService(storage storage.Goal, logger *slog.Logger) *GoalService {
	return &GoalService{
		Goal:   storage,
		logger: logger,
	}
}

func (b *GoalService) CreateGoal(ctx context.Context, req *pb.CreateGoalReq) (*pb.GoalResp, error) {
	resp, err := b.Goal.CreateGoal(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not create goal: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) GetGoal(ctx context.Context, req *pb.GetGoalReq) (*pb.GoalResp, error) {
	resp, err := b.Goal.GetGoal(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not get goal: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) UpdateGoal(ctx context.Context, req *pb.UpdateGoalReq) (*pb.GoalResp, error) {
	resp, err := b.Goal.UpdateGoal(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not update goal: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) DeleteGoal(ctx context.Context, req *pb.DeleteGoalReq) (*pb.Void3, error) {
	_, err := b.Goal.DeleteGoal(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not delete goal: %v", err))
		return nil, err
	}
	return &pb.Void3{}, nil
}
func (b *GoalService) ListGoals(ctx context.Context, req *pb.ListGoalsReq) (*pb.ListGoalsResp, error) {
	resp, err := b.Goal.ListGoals(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not list goals: %v", err))
		return nil, err
	}
	return resp, nil
}

func (b *GoalService) GetUserSpending(ctx context.Context, rep *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	resp, err := b.Goal.GetUserSpending(rep)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not generate report: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) GetUserIncome(ctx context.Context, rep *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	resp, err := b.Goal.GetUserIncome(rep)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not generate report: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) GetGoalReportProgress(ctx context.Context, rep *pb.GoalProgressRequest) (*pb.GoalProgressResponse, error) {
	resp, err := b.Goal.GetGoalReportProgress(rep)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not generate report: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *GoalService) GetBudgetSummary(ctx context.Context, rep *pb.UserId) (*pb.GetUserBudgetResponse, error) {
	resp, err := b.Goal.GetBudgetSummary(rep)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not generate report: %v", err))
		return nil, err
	}
	return resp, nil
}
