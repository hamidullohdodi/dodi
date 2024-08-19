package category

import (
	pb "budgeting/genproto/category"
	"budgeting/storage/mongodb"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewCategoryRepo(mdb, &log)
	req := &pb.CreateCategoryReq{
		Name:        "Guli",
		Description: "assdfdadas",
	}
	res, err := budgeting.CreateCategory(req)
	if err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}
	fmt.Println(res)
}

func TestUpdateCategory(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewCategoryRepo(mdb, &log)
	req := &pb.UpdateCategoryReq{
		Id:          "66ba96ffe260a336c66e1d42",
		Name:        "llll",
		Description: "asdadas",
	}
	res, err := budgeting.UpdateCategory(req)
	if err != nil {
		t.Fatalf("Failed to update category: %v", err)
	}
	fmt.Println(res)
}

func TestDeleteCategory(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewCategoryRepo(mdb, &log)
	req := &pb.DeleteCategoryReq{
		UserId: "66ba96ffe260a336c66e1d42",
	}
	res, err := budgeting.DeleteCategory(req)
	if err != nil {
		t.Fatalf("Failed to delete category: %v", err)
	}
	fmt.Println(res)
}

func TestListCategories(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}

	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	log := slog.New(slog.NewTextHandler(log.Writer(), nil))
	repo := NewCategoryRepo(mdb, log)

	req := &pb.ListCategoriesReq{
		Limit:  "10",
		Paid: "0",
	}

	res, err := repo.ListCategories(req)
	if err != nil {
		t.Fatalf("Failed to list budgets: %v", err)
	}

	for _, budget := range res.Categories {
		fmt.Printf("Budget: %+v\n", budget)
	}
	if len(res.Categories) == 0 {
		t.Errorf("Expected budgets but got none")
	}
}
