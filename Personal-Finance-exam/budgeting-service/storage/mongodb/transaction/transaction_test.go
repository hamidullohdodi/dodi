package transaction

import (
	pb "budgeting/genproto/transaction"
	"budgeting/storage/mongodb"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
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

	budgeting := NewTransactionRepo(mdb, &log)

	id := "66b92c7395c356f2ad272b6f"

	req := &pb.CreateTransactionReq{
		AccountId:   id,
		CategoryId:  id,
		Description: "asdadas",
	}
	res, err := budgeting.CreateTransaction(req)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}
	fmt.Println(res)
}

func TestGetTransaction(t *testing.T) {
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
	budgeting := NewTransactionRepo(mdb, &log)
	req := &pb.GetTransactionReq{
		UserId: "66ba971eb73e07396955e059",
	}
	res, err := budgeting.GetTransaction(req)
	if err != nil {
		t.Fatalf("Failed to get transaction: %v", err)
	}
	fmt.Println(res)
}

func TestUpdateTransaction(t *testing.T) {
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
	budgeting := NewTransactionRepo(mdb, &log)
	req := &pb.UpdateTransactionReq{
		Id:          "66ba971eb73e07396955e059",
		AccountId:   "66b93123bd921480d7a8803c",
		CategoryId:  "66b93123bd921480d7a8803c",
		Description: "asdadas",
		Amount:      "aaa",
		Date:        "aaa",
	}
	res, err := budgeting.UpdateTransaction(req)
	if err != nil {
		t.Fatalf("Failed to update transaction: %v", err)
	}
	fmt.Println(res)
}

func TestDeleteTransaction(t *testing.T) {
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
	budgeting := NewTransactionRepo(mdb, &log)
	req := &pb.DeleteTransactionReq{
		UserId: "66ba971eb73e07396955e059",
	}
	res, err := budgeting.DeleteTransaction(req)
	if err != nil {
		t.Fatalf("Failed to delete transaction: %v", err)
	}
	fmt.Println(res)
}

func TestListTransactions(t *testing.T) {
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
	repo := NewTransactionRepo(mdb, log)

	req := &pb.ListTransactionsReq{
		Limit:  "10",
		Offset: "0",
	}

	res, err := repo.ListTransactions(req)
	if err != nil {
		t.Fatalf("Failed to list budgets: %v", err)
	}

	for _, budget := range res.Transactions {
		fmt.Printf("Budget: %+v\n", budget)
	}
	if len(res.Transactions) == 0 {
		t.Errorf("Expected budgets but got none")
	}
}
