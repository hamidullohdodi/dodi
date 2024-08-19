package service

import (
	pb "budgeting/genproto/category"
	"budgeting/storage"
	"context"
	"fmt"
	"log/slog"
)

type CategoryIService struct {
	pb.UnimplementedCategoryServiceServer
	storage.CategoryI
	logger *slog.Logger
}

func NewCategoryIService(storage storage.CategoryI, logger *slog.Logger) *CategoryIService {
	return &CategoryIService{
		CategoryI: storage,
		logger:    logger,
	}
}

func (b *CategoryIService) CreateCategory(ctx context.Context, req *pb.CreateCategoryReq) (*pb.CategoryResp, error) {
	resp, err := b.CategoryI.CreateCategory(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not create category: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *CategoryIService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryReq) (*pb.CategoryResp, error) {
	resp, err := b.CategoryI.UpdateCategory(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not update category: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *CategoryIService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) (*pb.Void2, error) {
	_, err := b.CategoryI.DeleteCategory(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not delete category: %v", err))
		return nil, err
	}
	return &pb.Void2{}, nil
}
func (b *CategoryIService) ListCategories(ctx context.Context, req *pb.ListCategoriesReq) (*pb.ListCategoriesResp, error) {
	resp, err := b.CategoryI.ListCategories(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not list categories: %v", err))
		return nil, err
	}
	return resp, nil
}
