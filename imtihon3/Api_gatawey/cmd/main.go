package main

import (
	"api_getway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	product, err := grpc.NewClient("productserver:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	auth, err := grpc.NewClient("authservice:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	r := api.NewRouter(product, auth)

	r.Run("apigateway:50051")
}
