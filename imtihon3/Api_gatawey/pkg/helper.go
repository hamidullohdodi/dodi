package pkg

import (
	ap "api_getway/genproto"
	pb "api_getway/genproto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type Client1 struct {
	AuthClient    ap.AuthServiceClient
	ProductClient pb.ProductServiceClient
}

func NewClient() *Client1 {
	product, err1 := grpc.NewClient("productservers:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		slog.Error("Failed to connect to gRPC server", slog.String("address", "localhost:8080"), slog.String("error", err1.Error()))
	}
	auth, err2 := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		slog.Error("Failed to connect to gRPC server", slog.String("address", "localhost:8080"), slog.String("error", err2.Error()))
	}

	proC := pb.NewProductServiceClient(product)
	authC := ap.NewAuthServiceClient(auth)

	return &Client1{
		AuthClient:    authC,
		ProductClient: proC,
	}
}
