package budget

import (
	pb "budgeting/genproto/budget"
	"budgeting/storage/mongodb"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestCreateBudget(t *testing.T) {
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
	budgeting := NewBudgetRepo(mdb, &log)
	req := &pb.CreateBudgetReq{
		Name:      "Guli",
		Amount:    "aaaa",
		StartingTime: ".......",
		EndingTime:   ".......",
	}
	res, err := budgeting.CreateBudget(req)
	if err != nil {
		t.Fatalf("Failed to create budget: %v", err)
	}
	fmt.Println(res)
}

func TestUpdateBudget(t *testing.T) {
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
	budgeting := NewBudgetRepo(mdb, &log)
	req := &pb.UpdateBudgetReq{
		Id:        "66ba938e981514d139eec0af",
		Name:      "Guli",
		Amount:    "333",
		StartingTime: "111",
		EndingTime:   "222",
	}
	res, err := budgeting.UpdateBudget(req)
	if err != nil {
		t.Fatalf("Failed to update budget: %v", err)
	}
	fmt.Println(res)
}

func TestDeleteBudget(t *testing.T) {
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
	budgeting := NewBudgetRepo(mdb, &log)
	req := &pb.DeleteBudgetReq{
		UserId: "66ba938e981514d139eec0af",
	}
	res, err := budgeting.DeleteBudget(req)
	if err != nil {
		t.Fatalf("Failed to delete budget: %v", err)
	}
	fmt.Println(res)
}

func TestListBudgets(t *testing.T) {
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
	repo := NewBudgetRepo(mdb, log)

	req := &pb.ListBudgetsReq{
		Limit:  "10",
		Paid: "0",
	}

	res, err := repo.ListBudgets(req)
	if err != nil {
		t.Fatalf("Failed to list budgets: %v", err)
	}

	for _, budget := range res.Budgets {
		fmt.Printf("Budget: %+v\n", budget)
	}
	if len(res.Budgets) == 0 {
		t.Errorf("Expected budgets but got none")
	}
}
