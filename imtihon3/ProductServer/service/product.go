package service

import (
	"context"
	"database/sql"
	pb "product_service/genproto/product"
	"product_service/storage/postgres"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	Repo *postgres.ProductRepo
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		Repo: postgres.NewProductRepo(db),
	}
}

func (p *ProductService) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.ProductResponse, error) {
	r, err := p.Repo.AddProduct(req)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return r, nil
}

func (p *ProductService) EditProduct(ctx context.Context, req *pb.EditProductRequest) (*pb.ProductResponse, error) {
	r, err := p.Repo.EditProduct(req)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return r, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	r, err := p.Repo.DeleteProduct(req)
	if err != nil {
		return &pb.DeleteProductResponse{}, err
	}
	return r, nil
}
func (p *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	r, err := p.Repo.GetProduct(req)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return r, nil
}
func (p *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	r, err := p.Repo.ListProducts(req)
	if err != nil {
		return &pb.ListProductsResponse{}, err
	}
	return r, nil
}

func (p *ProductService) AddRating(ctx context.Context, req *pb.AddRatingRequest) (*pb.RatingResponse, error) {
	r, err := p.Repo.AddRating(req)
	if err != nil {
		return &pb.RatingResponse{}, err
	}
	return r, nil
}
func (p *ProductService) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.ListProductsResponse, error) {
	r, err := p.Repo.SearchProducts(req)
	if err != nil {
		return &pb.ListProductsResponse{}, err
	}
	return r, nil
}

func (p *ProductService) ListRatings(ctx context.Context, req *pb.ListRatingsRequest) (*pb.ListRatingsResponse, error) {
	r, err := p.Repo.ListRatings(req)
	if err != nil {
		return &pb.ListRatingsResponse{}, err
	}
	return r, nil
}
