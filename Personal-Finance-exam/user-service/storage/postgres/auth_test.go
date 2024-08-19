package postgres

import (
	"user/config"
	pb "user/genproto/auth"
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestRegister(t *testing.T) {

	cfg, err := config.Load(".")
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	conn := fmt.Sprintf("host=postgresdb user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Name,
		cfg.Database.Password,
		cfg.Database.Port,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	auth := NewAuthRepo(db)

	req := pb.RegisterUserReq{
		FirstName:    "Aliy",
		Email:       "aliy@gmail.com1",
		Password:    "aliy",
		LastName:     "Aliyev",
		Role:        "admin",
		DateOfBirth: "10.06.2006",
	}
	resp, err := auth.Register(&req)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}
	fmt.Println(resp)
}

func TestLogin(t *testing.T) {
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

	cfg, err := config.Load(".")
	if err != nil{
		fmt.Println(err)
	}

	conn := fmt.Sprintf("host=postgresdb user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Name,
		cfg.Database.Password,
		cfg.Database.Port,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		t.Fatal(err)
	}

	auth := NewAuthRepo(db)

	req := pb.LoginUserReq{
		FirstName: "Ali",
		Password: "aliy",
	}

	resp, err := auth.Login(&req)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	fmt.Println(resp)
}
