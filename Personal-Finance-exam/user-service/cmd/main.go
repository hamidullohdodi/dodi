package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"

	"user/api"
	_ "user/api/docs"
	"user/api/handler"
	"user/config"
	user "user/genproto/user"
	"user/service"
	"user/storage/postgres"
)

func main() {
	cfg, err := config.Load(".")
	if err != nil {
		fmt.Println(err)
	}

	db, err := postgres.NewPostgresStorage(*cfg)
	if err != nil {
		log.Fatalf("can't connect to db: %v", err)
	}
	defer db.Db.Close()

	targetListen := fmt.Sprintf("%s:%s", cfg.AuthHost, cfg.AuthPort)
	listener, err := net.Listen("tcp", targetListen)
	fmt.Println(targetListen)
	if err != nil {
		slog.Error("can't listen: %v", err)
	}

	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, userService)

	go func() {
		slog.Info("gRPC server started on port %s", cfg.AuthPort)
		if err := s.Serve(listener); err != nil {
			slog.Error("can't serve: %v", err)
		}
	}()

	h := handler.NewHandler(authService)
	router := api.Engine(h)

	log.Printf("UserSerivce Running on :%s port", cfg.UserPort)
	target := fmt.Sprintf("%s:%s", cfg.UserHost, cfg.UserPort)
	if err := router.Run(target); err != nil {
		log.Fatalf("can't start server: %v", err)
	}

	log.Printf("REST server started on port %s", cfg.UserPort)
}
