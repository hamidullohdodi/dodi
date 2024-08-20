package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	pb "product_service/genproto/product"
	"time"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

//func (o *OrderRepo) CreateOrder(rep *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
//	orderID := uuid.New().String()
//
//	var totalAmount float64
//	itemsResponse := make([]*pb.OrderItem, 0)
//
//	for _, item := range rep.Items {
//		var price float64 = 0
//
//		itemResponse := &pb.OrderItem{
//			ProductId: item.ProductId,
//			Quantity:  item.Quantity,
//			Price:     price,
//		}
//
//		totalAmount += price * float64(item.Quantity)
//
//		itemsResponse = append(itemsResponse, itemResponse)
//	}
//
//	shippingAddressResponse := &pb.ShippingAddress{
//		Street:  rep.ShippingAddress.Street,
//		City:    rep.ShippingAddress.City,
//		Country: rep.ShippingAddress.Country,
//		ZipCode: rep.ShippingAddress.ZipCode,
//	}
//
//	orderResponse := &pb.OrderResponse{
//		Id:              orderID,
//		UserId:          "user789",
//		Items:           itemsResponse,
//		TotalAmount:     totalAmount,
//		Status:          "pending",
//		ShippingAddress: shippingAddressResponse,
//		CreatedAt:       time.Now().Format(time.RFC3339),
//	}
//
//	_, err := o.db.Exec(`
//		INSERT INTO orders (id, user_id, total_amount, status, shipping_address, created_at)
//		VALUES ($1, $2, $3, $4, $5, $6)`,
//		orderResponse.Id, orderResponse.UserId, orderResponse.TotalAmount, orderResponse.Status, shippingAddressResponse.String(), orderResponse.CreatedAt,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	return orderResponse, nil
//}

func (o *OrderRepo) CancelOrder(req *pb.CancelOrderRequest) (*pb.UpdateOrderStatusResponse, error) {
	_, err := uuid.Parse(req.OrderId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.UpdateOrderStatusResponse{}, err
	}

	_, err = o.db.Exec("UPDATE orders SET status = 'cancelled', updated_at = $1 WHERE id = $2", time.Now(), req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderStatusResponse{
		OrderId:   req.OrderId,
		Status:    "cancelled",
		UpdatedAt: time.Now().String(),
	}, nil
}

func (o *OrderRepo) ListOrders(req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	offset := (req.Page - 1) * req.Limit
	rows, err := o.db.Query("SELECT id, user_id, total_amount, status, created_at FROM orders LIMIT $1 OFFSET $2", req.Limit, offset)
	if err != nil {
		return nil, err
	}

	orders := []*pb.OrderResponse{}
	for rows.Next() {
		var order pb.OrderResponse
		err := rows.Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	var total int32
	err = o.db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return nil, err
	}

	response := &pb.ListOrdersResponse{
		Orders: orders,
		Total:  total,
		Page:   req.Page,
		Limit:  req.Limit,
	}
	return response, nil
}

func (o *OrderRepo) UpdateOrderStatus(req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	_, err := uuid.Parse(req.OrderId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.UpdateOrderStatusResponse{}, err
	}

	_, err = o.db.Exec("UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3",
		req.Status, time.Now(), req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderStatusResponse{
		OrderId:   req.OrderId,
		Status:    req.Status,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (o *OrderRepo) GetOrder(req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	_, err := uuid.Parse(req.OrderId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.OrderResponse{}, err
	}
	order := &pb.OrderResponse{}
	err = o.db.QueryRow("SELECT user_id, total_amount, status, created_at FROM orders where id = $1", req.OrderId).Scan(&order.UserId, &order.TotalAmount, &order.Status, &order.CreatedAt)
	order.Id = req.OrderId
	if err != nil {
		return &pb.OrderResponse{}, err
	}
	return order, nil
}

func (o *OrderRepo) UpdateShipping(req *pb.UpdateShippingRequest) (*pb.ShippingResponse, error) {
	_, err := uuid.Parse(req.OrderId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.ShippingResponse{}, err
	}
	updatedAt := time.Now().String()
	_, err = o.db.Exec("UPDATE shipping_info SET tracking_number = $1, carrier = $2, estimated_delivery_date = $3, updated_at = $4 WHERE id = $5",
		req.TrackingNumber, req.Carrier, req.EstimatedDeliveryDate, updatedAt, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &pb.ShippingResponse{
		OrderId:               req.OrderId,
		TrackingNumber:        req.TrackingNumber,
		Carrier:               req.Carrier,
		EstimatedDeliveryDate: req.EstimatedDeliveryDate,
		UpdatedAt:             updatedAt,
	}, nil
}

func (o *OrderRepo) OrderPayment(req *pb.OrderPaymentRequest) (*pb.PaymentResponse, error) {
	req.OrderId = uuid.New().String()
	create := time.Now().String()
	_, err := o.db.Exec(`
		INSERT INTO order_payments(order_id, payment_method, card_number, expiry_date, cvv, created_at)
		VALUES($1, $2, $3, $4, $5, $6)`,
		req.OrderId, req.PaymentMethod, req.CardNumber, req.ExpiryDate, req.Cvv, create)

	if err != nil {
		return nil, err
	}
	amount := 599.97

	resp := &pb.PaymentResponse{
		OrderId:       req.OrderId,
		PaymentId:     uuid.New().String(),
		Amount:        amount,
		Status:        "success",
		TransactionId: "tx789",
		CreatedAt:     create,
	}

	return resp, nil
}
func (o *OrderRepo) CheckPaymentStatus(rep *pb.CheckPaymentStatusRequest) (*pb.PaymentResponse, error) {
	_, err := uuid.Parse(rep.OrderId)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return nil, err
	}

	order := &pb.PaymentResponse{}
	err = o.db.QueryRow("SELECT order_id, payment_method FROM order_payments WHERE order_id = $1", rep.OrderId).Scan(
		&order.OrderId, &order.PaymentMethod)
	if err != nil {
		return nil, err
	}

	return order, nil
}
