package service

import (
	"context"
	"database/sql"
	pb "product_service/genproto/product"
	"product_service/storage/postgres"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	Repo *postgres.OrderRepo
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		Repo: postgres.NewOrderRepo(db),
	}
}

//func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
//	r, err := o.Repo.CreateOrder(req)
//	if err != nil {
//		return &pb.OrderResponse{}, err
//	}
//	return r, nil
//}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	r, err := o.Repo.UpdateOrderStatus(req)
	if err != nil {
		return &pb.UpdateOrderStatusResponse{}, err
	}
	return r, nil
}

func (o *OrderService) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.UpdateOrderStatusResponse, error) {
	r, err := o.Repo.CancelOrder(req)
	if err != nil {
		return &pb.UpdateOrderStatusResponse{}, err
	}
	return r, nil
}

func (o *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	r, err := o.Repo.ListOrders(req)
	if err != nil {
		return &pb.ListOrdersResponse{}, err
	}
	return r, nil
}

func (o *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	r, err := o.Repo.GetOrder(req)
	if err != nil {
		return &pb.OrderResponse{}, err
	}
	return r, nil
}

func (o *OrderService) UpdateShipping(ctx context.Context, req *pb.UpdateShippingRequest) (*pb.ShippingResponse, error) {
	r, err := o.Repo.UpdateShipping(req)
	if err != nil {
		return &pb.ShippingResponse{}, err
	}
	return r, nil
}

func (o *OrderService) OrderPayment(ctx context.Context, req *pb.OrderPaymentRequest) (*pb.PaymentResponse, error) {
	r, err := o.Repo.OrderPayment(req)
	if err != nil {
		return &pb.PaymentResponse{}, err
	}
	return r, nil
}

func (o *OrderService) CheckPaymentStatus(ctx context.Context, req *pb.CheckPaymentStatusRequest) (*pb.PaymentResponse, error) {
	r, err := o.Repo.CheckPaymentStatus(req)
	if err != nil {
		return &pb.PaymentResponse{}, err
	}
	return r, nil
}
