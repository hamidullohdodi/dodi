package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	pb "product_service/genproto/product"
	"product_service/service"
	"product_service/storage/postgres"
)

func main() {
	logger := slog.Logger{}

	db, err := postgres.ConnectDB()
	if err != nil {
		logger.Error(err.Error())
	}
	listener, err := net.Listen("tcp", "productserver:8082")
	if err != nil {
		logger.Error(err.Error())
	}
	server := grpc.NewServer()

	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	pb.RegisterProductServiceServer(server, service.NewProductService(db))
	pb.RegisterOrderServiceServer(server, service.NewOrderService(db))

	log.Println("server listned at :8082")

	err = server.Serve(listener)
	if err != nil {
		logger.Error(err.Error())
	}
}
