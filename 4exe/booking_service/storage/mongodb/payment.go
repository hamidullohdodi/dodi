package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *BookingRepo) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	collection := p.db.Collection("payments")

	// To'lovni yaratish
	payment := bson.M{
		"booking_id":     req.GetBookingId(),
		"amount":         req.GetAmount(),
		"status":         req.GetStatus(),
		"payment_method": req.GetPaymentMethod(),
		"transaction_id": req.GetTransactionId(),
	}

	// To'lovni ma'lumotlar bazasiga saqlash
	resp, err := collection.InsertOne(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %v", err)
	}

	// ID'ni olish va formatlash
	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID: %v", resp.InsertedID)
	}

	// Javobni tayyorlash
	return &pb.CreatePaymentResponse{
		BookingId:     req.GetBookingId(),
		Amount:        req.GetAmount(),
		Status:        req.GetStatus(),
		PaymentMethod: req.GetPaymentMethod(),
		TransactionId: req.GetTransactionId(),
		Id:            oid.Hex(),
	}, nil
}

func (p *BookingRepo) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	collection := p.db.Collection("payments")

	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("provided ID is not a valid ObjectID: %v", err)
	}

	filter := bson.M{"_id": oid}

	var payment struct {
		BookingId     string  `bson:"booking_id"`
		Amount        float64 `bson:"amount"` // Typeni `float64` qilib o'zgartirish
		Status        string  `bson:"status"`
		PaymentMethod string  `bson:"payment_method"`
		TransactionID string  `bson:"transaction_id"`
		XId           string  `bson:"_id"`
	}

	err = collection.FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		return nil, fmt.Errorf("failed to find payment: %v", err)
	}

	resp := pb.GetPaymentResponse{
		BookingId:     payment.BookingId,
		Status:        payment.Status,
		PaymentMethod: payment.PaymentMethod,
		TransactionId: payment.TransactionID,
		Id:            payment.XId,
		Amount:        float32(payment.Amount), // Floatni stringga o'girish
	}

	return &resp, nil
}



func (p *BookingRepo) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
	collection := p.db.Collection("payments")

	// Limit and offset for pagination
	limit := int64(req.GetLimit())
	offset := int64(req.GetOffset())

	// Query options
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	// Finding payments
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find payments: %v", err)
	}
	defer cursor.Close(ctx)

	// Preparing the response
	var payments []*pb.Payment
	for cursor.Next(ctx) {
		var payment struct {
			BookingId     string  `bson:"booking_id"`
			Amount        float64 `bson:"amount"`
			Status        string  `bson:"status"`
			PaymentMethod string  `bson:"payment_method"`
			TransactionID string  `bson:"transaction_id"`
			Id           string  `bson:"_id"`
		}

		err := cursor.Decode(&payment)
		if err != nil {
			return nil, fmt.Errorf("failed to decode payment: %v", err)
		}

		payments = append(payments, &pb.Payment{
			BookingId:     payment.BookingId,
			Amount:        float32(payment.Amount),
			Status:        payment.Status,
			PaymentMethod: payment.PaymentMethod,
			TransactionId: payment.TransactionID,
			Id:            payment.Id,
		})
	}

	// Check for any errors encountered during iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.ListPaymentsResponse{
		Payments: payments,
	}, nil
}
func (p *BookingRepo) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error) {
	collection := p.db.Collection("payments")

	// Converting the ID from the request to ObjectID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid payment ID: %v", err)
	}

	// Creating the update document
	update := bson.M{
		"$set": bson.M{
			"booking_id":     req.GetBookingId(),
			"amount":         req.GetAmount(),
			"status":         req.GetStatus(),
			"payment_method": req.GetPaymentMethod(),
			"transaction_id": req.GetTransactionId(),
		},
	}

	// Defining the filter to find the specific payment by ID
	filter := bson.M{"_id": oid}

	// Updating the payment document
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update payment: %v", err)
	}

	// Checking if any document was modified
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no payment found with the given ID")
	}

	// Preparing the response

	return &pb.UpdatePaymentResponse{}, nil
}

func (p *BookingRepo) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	collection := p.db.Collection("payments")

	// Convert the payment ID from the request to ObjectID
	log.Print(req.Id)
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid payment ID: %v", err)
	}

	// Define the filter to find the specific payment by ID
	filter := bson.M{"_id": oid}

	// Delete the payment document
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete payment: %v", err)
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no payment found with the given ID")
	}

	return &pb.DeletePaymentResponse{Message: "Payment deleted successfully"}, nil
}



