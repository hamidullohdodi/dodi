package main

import (
	"auth_service/api"
	"auth_service/api/handler"
	"auth_service/config"
	"auth_service/genproto/user"
	"auth_service/logger"
	"auth_service/service"
	"auth_service/storage/postgres"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {

	Db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	log := logger.NewLogger()

	userRepo := postgres.NewAuthService(Db)
	userService := service.NewAuthService(userRepo, log)
	userHandler := handler.NewAuthenticaionHandlerImpl(userService, log)

	x := grpc.NewServer()

	user.RegisterAuthServer(x, userService)

	var cat sync.WaitGroup
	cat.Add(1)

	go func() {
		fmt.Printf("start gRPC server on port %s\n", config.Load().USER_SERVICE)
		auth := api.NewApiService(userHandler)
		router := auth.Router()
		router.Run(config.Load().USER_SERVICE)
		cat.Done()
	}()

	list, err := net.Listen("tcp", "auth:50051")

	if err != nil {
		log.Error("Error: %v", err)
		return
	}
	if err := x.Serve(list); err != nil {
		log.Error("Error: %v", err)
		return
	}

	cat.Wait()

}
