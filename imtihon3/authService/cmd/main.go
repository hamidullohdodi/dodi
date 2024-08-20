package main

import (
	pb "auth_service/genproto"
	"auth_service/service"
	"auth_service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	liss, err := net.Listen("tcp", "authservice:8083")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, service.NewUserService(db))

	log.Printf("server listening at %v", liss.Addr())

	if err := server.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
