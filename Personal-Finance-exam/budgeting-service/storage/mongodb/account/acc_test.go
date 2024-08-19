package account

import (
	pb "budgeting/genproto/account"
	"budgeting/storage/mongodb"
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestCreateAccount(t *testing.T) {
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

	budgeting := NewAccountRepo(mdb, &log)

	req := &pb.CreateAccountReq{
		UserId: uuid.New().String(),
		Name:        "Muslima",
		Description: "...",
	}
	res, err := budgeting.CreateAccount(req)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}
	fmt.Println(res)
}

func TestGetAccount(t *testing.T) {
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

	budgeting := NewAccountRepo(mdb, &log)

	id := "2a7af15c-8eac-4172-8f9d-4dc5c7dc748f"

	req := &pb.GetAccountReq{
		UserId: id,
	}

	res, err := budgeting.GetAccount(req)
	if err != nil {
		t.Fatalf("Failed to get account: %v", err)
	}
	fmt.Println(res)
}

func TestUpdateAccount(t *testing.T) {
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

	budgeting := NewAccountRepo(mdb, &log)

	id := "66ba63447afafa16345736ed"

	req := &pb.UpdateAccountReq{
		Id:          id,
		Name:        "hamidulloh",
		Description: "asdadas",
	}
	res, err := budgeting.UpdateAccount(req)
	if err != nil {
		t.Fatalf("Failed to update account: %v", err)
	}
	fmt.Println(res)
}

func TestDeleteAccount(t *testing.T) {
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

	budgeting := NewAccountRepo(mdb, &log)

	id := "66ba63447afafa16345736ed"

	req := &pb.DeleteAccountReq{
		UserId: id,
	}
	res, err := budgeting.DeleteAccount(req)
	if err != nil {
		t.Fatalf("Failed to delete account: %v", err)
	}
	fmt.Println(res)
}

func TestListAccount(t *testing.T) {
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
	budgeting := NewAccountRepo(mdb, &log)

	rep := &pb.ListAccountsReq{
		Limit:  "2",
		Paid: "0",
	}

	res, err := budgeting.ListAccounts(rep)
	if err != nil {
		t.Fatalf("Failed to list accounts: %v", err)
	}

	fmt.Println("Response:", res)
}
