package postgres

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	pb "product_service/genproto/product"
	"testing"
	"time"
)

//func TestCreateOrder(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	orderRepo := NewOrderRepo(db)
//
//	OrderItem := &pb.OrderItem{
//		ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
//		Quantity:  10,
//		Price: 12.12,
//    }
//
//	mockCreateRequest := &pb.CreateOrderRequest{
//		pb.OrderItem{
//			ProductId: "4317b534-0cdc-408a-b43c-86ebee331281",
//
//		}
//		ShippingAddress: &pb.ShippingAddress{
//			City:    "City",
//			Street:  "Street",
//			ZipCode: "123",
//			Country: "Country",
//		},
//	}
//	_, err = orderRepo.CreateOrder(mockCreateRequest)
//
//}

func TestUpdateOrderStatus(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	mockUpdateRequest := &pb.UpdateOrderStatusRequest{
		OrderId: "157ed58a-4dc7-470c-9c93-52dcdd11f773",
		Status:  "cancelled",
	}

	_, err = orderRepo.UpdateOrderStatus(mockUpdateRequest)
	if err != nil {
		t.Errorf("Error updating order status: %v", err)
	}

}

func TestCancelOrder(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	cancelOrderReq := &pb.CancelOrderRequest{
		OrderId: "157ed58a-4dc7-470c-9c93-52dcdd11f773",
	}

	response, err := orderRepo.CancelOrder(cancelOrderReq)
	if err != nil {
		t.Errorf("Error canceling order: %v", err)
	}
	assert.Equal(t, response.Status, "cancelled")

}

func TestListOrders(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	listOrdersReq := &pb.ListOrdersRequest{
		Page:  3,
		Limit: 4,
	}

	response, err := orderRepo.ListOrders(listOrdersReq)
	if err != nil {
		t.Errorf("Error listing orders: %v", err)
	}
	assert.Equal(t, len(response.Orders), 0)

}

func TestGetOrder(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	getOrderReq := &pb.GetOrderRequest{
		OrderId: "157ed58a-4dc7-470c-9c93-52dcdd11f773",
	}

	response, err := orderRepo.GetOrder(getOrderReq)
	if err != nil {
		t.Errorf("Error getting orders: %v", err)
	}
	fmt.Printf("%+v\n", response)

}

func TestUpdateShipping(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	updateShippingReq := &pb.UpdateShippingRequest{
		OrderId:               "157ed58a-4dc7-470c-9c93-52dcdd11f773",
		TrackingNumber:        "123456789",
		Carrier:               "UPS",
		EstimatedDeliveryDate: time.Now().String(),
	}
	_, err = orderRepo.UpdateShipping(updateShippingReq)
	if err != nil {
		t.Errorf("Error updating shipping: %v", err)
	}

}

func TestOrderPayment(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}

	payOrderReq := &pb.OrderPaymentRequest{
		OrderId:       "157ed58a-4dc7-470c-9c93-52dcdd11f773",
		PaymentMethod: "1121232",
		ExpiryDate:    "12/27",
		Cvv:           "123",
	}
	_, err = orderRepo.OrderPayment(payOrderReq)
	resp, err := orderRepo.OrderPayment(payOrderReq)
	if err != nil {
		t.Errorf("Error ordering payment: %v", err)
	}

	if resp.OrderId == "" || resp.PaymentId == "" || resp.Status != "success" || resp.Amount != 599.97 {
		t.Errorf("Unexpected response: %+v", resp)
	}
}

func TestCheckPaymentStatus(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	orderRepo := &OrderRepo{db: db}
	checkPaymentStatusReq := &pb.CheckPaymentStatusRequest{
		OrderId: "467e79ea-30fe-4d6e-89dd-b6e33b91cb69",
	}
	checkPaymentStatusResp, err := orderRepo.CheckPaymentStatus(checkPaymentStatusReq)
	if err != nil {
		t.Errorf("Error checking payment status: %v", err)
	}
	fmt.Printf("%+v\n", checkPaymentStatusResp)

}
