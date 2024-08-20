package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"log"
	pb "product_service/genproto/product"
	"time"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) AddProduct(req *pb.AddProductRequest) (*pb.ProductResponse, error) {
	var product pb.ProductResponse
	ruuid := uuid.New()
	createdAt := time.Now()

	_, err := p.db.Exec("INSERT INTO products(id, name, description, price, category_id, quantity, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		ruuid, req.Name, req.Description, req.Price, req.CategoryId, req.Quantity, createdAt.String())
	if err != nil {
		return nil, err
	}
	product.Id = ruuid.String()
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.CategoryId = req.CategoryId
	product.Quantity = req.Quantity
	product.CreatedAt = createdAt.String()

	return &product, nil
}

func (p *ProductRepo) EditProduct(req *pb.EditProductRequest) (*pb.ProductResponse, error) {
	_, err := uuid.Parse(req.ProductId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.ProductResponse{}, err
	}

	_, err = p.db.Exec("UPDATE products SET name=$1, description=$2, price=$3, category_id=$4, quantity=$5, updated_at=$6 WHERE id=$7",
		req.Name, req.Description, req.Price, req.CategoryId, req.Quantity, time.Now().String(), req.ProductId)
	if err != nil {
		return nil, err
	}

	var product pb.ProductResponse
	err = p.db.QueryRow("SELECT id, name, description, price, category_id, quantity, created_at, updated_at FROM products WHERE id=$1", req.ProductId).
		Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepo) DeleteProduct(req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	_, err := uuid.Parse(req.ProductId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.DeleteProductResponse{}, err
	}
	_, err = p.db.Exec("update payments set deleted_at = EXTRACT(EPOCH FROM NOW()) where id = $1", req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{
		Message: "Product successfully deleted",
	}, nil
}

func (p *ProductRepo) ListProducts(req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	rows, err := p.db.Query(
		"SELECT id, name, price, category_id FROM products LIMIT $1 OFFSET $2;",
		req.Limit, (req.Page-1)*req.Limit,
	)
	if err != nil {
		return nil, err
	}

	var products []*pb.ProductResponse
	for rows.Next() {
		var product pb.ProductResponse
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	var total int32
	err = p.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		return nil, err
	}

	return &pb.ListProductsResponse{
		Products: products,
		Total:    total,
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (p *ProductRepo) GetProduct(req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	_, err := uuid.Parse(req.ProductId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.ProductResponse{}, err
	}
	var product pb.ProductResponse

	err = p.db.QueryRow(
		"SELECT id, name, description, price, category_id, quantity FROM products WHERE id=$1",
		req.ProductId,
	).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Quantity)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepo) SearchProducts(req *pb.SearchProductsRequest) (*pb.ListProductsResponse, error) {
	query := "SELECT id, name, price, category_id FROM products WHERE deleted_at = $1"
	var params []interface{}

	if req.Category != "" {
		query += " AND category_id = $2"
		params = append(params, req.Category)
	}

	if req.MinPrice > 0 {
		query += " AND price >= $3"
		params = append(params, req.MinPrice)
	}

	if req.MaxPrice > 0 {
		query += " AND price <= $4"
		params = append(params, req.MaxPrice)
	}

	query += " LIMIT $5, OFFSET $6"

	params = append(params, req.Limit, (req.Page-1)*req.Limit)

	rows, err := p.db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	var products []*pb.ProductResponse
	for rows.Next() {
		var product pb.ProductResponse
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	var total int32
	err = p.db.QueryRow("SELECT COUNT(*) FROM products ", params...).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &pb.ListProductsResponse{
		Products: products,
		Total:    total,
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (p *ProductRepo) AddRating(req *pb.AddRatingRequest) (*pb.RatingResponse, error) {
	var rating pb.RatingResponse

	err := p.db.QueryRow(
		"INSERT INTO ratings (product_id, rating, comment, created_at) VALUES ($1, $2, $3, $4 ) RETURNING id, product_id, user_id, rating, comment, created_at",
		req.ProductId, req.Rating, req.Comment, time.Now(),
	).Scan(&rating.Id, &rating.ProductId, &rating.UserId, &rating.Rating, &rating.Comment, &rating.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (p *ProductRepo) ListRatings(req *pb.ListRatingsRequest) (*pb.ListRatingsResponse, error) {
	rows, err := p.db.Query(`
		SELECT id, product_id, user_id, rating, comment, created_at FROM ratings WHERE product_id = $1`,
		req.ProductId,
	)
	if err != nil {
		return nil, err
	}

	var ratings []*pb.RatingResponse
	var totalRatings int32
	var totalRatingSum float64

	for rows.Next() {
		var rating pb.RatingResponse
		err := rows.Scan(&rating.Id, &rating.ProductId, &rating.UserId, &rating.Rating, &rating.Comment, &rating.CreatedAt)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
		totalRatings++
		totalRatingSum += rating.Rating
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var averageRating float64
	if totalRatings > 0 {
		averageRating = totalRatingSum / float64(totalRatings)
	}

	response := &pb.ListRatingsResponse{
		Ratings:       ratings,
		AverageRating: averageRating,
		TotalRatings:  totalRatings,
	}

	return response, nil
}

func (s *ProductRepo) CreateArtisanCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	id := uuid.New().String()
	createdAt := time.Now()

	query := `INSERT INTO artisan_categories (id, name, description, created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, id, req.Name, req.Description, createdAt)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   &pb.Timestamp{Seconds: createdAt.String()},
	}, nil
}

func (s *ProductRepo) CreateProductCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	id := uuid.New().String()
	createdAt := time.Now()

	query := `INSERT INTO product_categories (id, name, description, created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, id, req.Name, req.Description, createdAt)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   &pb.Timestamp{Seconds: createdAt.String()},
	}, nil
}
