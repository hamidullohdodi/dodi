package postgres

import (
	"log"
	pb "product_service/genproto/product"
	"testing"
)

func TestAddProduct(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	req := &pb.AddProductRequest{
		Name:        "Handmade",
		Description: "Beautiful handcrafted wooden chair",
		Price:       10.2,
		CategoryId:  "f00bb295-c4f5-4fc8-88d7-c309a6b45573",
		Quantity:    5,
	}

	resp, err := productRepo.AddProduct(req)

	if err != nil {
		t.Fatalf("AddProduct failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Expected non-nil ProductResponse, got nil")
	}
}
func TestEditProduct(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	editProductReq := &pb.EditProductRequest{
		ProductId:   "4317b534-0cdc-408a-b43c-86ebee331281",
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       129.99,
		CategoryId:  "f00bb295-c4f5-4fc8-88d7-c309a6b45573",
		Quantity:    15,
	}
	_, err = productRepo.EditProduct(editProductReq)
	if err != nil {
		t.Fatalf("EditProduct failed: %v", err)
	}
}

func TestDeleteProduct(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	deleteProductReq := &pb.DeleteProductRequest{
		ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
	}

	_, err = productRepo.DeleteProduct(deleteProductReq)
	if err != nil {
		t.Fatalf("DeleteProduct failed: %v", err)
	}

}

func TestListProducts(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	listProductsReq := &pb.ListProductsRequest{
		Page:  1,
		Limit: 3,
	}

	listProductsResp, err := productRepo.ListProducts(listProductsReq)
	if err != nil {
		t.Fatalf("ListProducts failed: %+v\n", err)
	}
	for _, product := range listProductsResp.Products {
		log.Printf("Product: %+v", product)
	}

}

func TestGetProduct(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	productRepo = NewProductRepo(db)

	getProductReq := &pb.GetProductRequest{
		ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
	}

	getProductResp, err := productRepo.GetProduct(getProductReq)
	if err != nil {
		t.Fatalf("GetProduct failed: %v", err)
	}
	log.Printf("Product: %v", getProductResp)
}
func TestSearchProducts(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	productRepo = NewProductRepo(db)

	req := &pb.SearchProductsRequest{
		Category: "1",
		MinPrice: 10,
		MaxPrice: 20,
		Page:     1,
		Limit:    3,
	}

	resp, err := productRepo.SearchProducts(req)
	if err != nil {
		t.Fatalf("SearchProducts failed: %v", err)
	}
	log.Printf("Products: %v", resp)
}

func TestAddRating(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	productRepo = NewProductRepo(db)

	rep := &pb.AddRatingRequest{
		ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
		Rating:    23.3,
		Comment:   "This is a comment",
	}

	_, err = productRepo.AddRating(rep)

}

func TestListRatings(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	productRepo := &ProductRepo{db: db}

	productRepo = NewProductRepo(db)

	listRatingsReq := &pb.ListRatingsRequest{
		ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
	}

	listRatingsResp, err := productRepo.ListRatings(listRatingsReq)
	if err != nil {
		t.Fatalf("ListRatings failed: %v", err)
	}
	log.Printf("Ratings: %v", listRatingsResp)
}
