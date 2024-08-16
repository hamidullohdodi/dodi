package pkg

import (
	"api-geteway/genproto/user"
	"api-geteway/genproto/booking"
	"api-geteway/internal/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthServiceClient(cfg *config.Config) user.AuthClient {
	conn, err := grpc.NewClient(cfg.USER, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return user.NewAuthClient(conn)
}

func NewBookingServiceClient(cfg *config.Config) booking.BookingClient {
	conn, err := grpc.NewClient(cfg.BOOKING, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return booking.NewBookingClient(conn)
}
