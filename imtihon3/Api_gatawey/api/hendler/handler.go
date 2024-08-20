package hendler

import (
	pbp "api_getway/genproto"
	pb "api_getway/genproto/product"
	"google.golang.org/grpc"
)

type Handler struct {
	Product pb.ProductServiceClient
	Auth    pbp.AuthServiceClient
}

func NewHandler(product, auth *grpc.ClientConn) *Handler {
	return &Handler{
		Product: pb.NewProductServiceClient(product),
		Auth:    pbp.NewAuthServiceClient(auth),
	}
}
