package pkg

import (
	"booking_service/config"
	user "booking_service/genproto/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthServiceClient(cfg *config.Config) user.AuthClient {
	conn, err := grpc.NewClient(cfg.BookingService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return user.NewAuthClient(conn)
}
