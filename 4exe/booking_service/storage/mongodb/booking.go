package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookingRepo struct {
	db *mongo.Database
}

func NewBookingRepo(db *mongo.Database) *BookingRepo {
	return &BookingRepo{db: db}
}

func (b *BookingRepo) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	collection := b.db.Collection("booking")

	booking := bson.M{
		"user_id":     req.UserId,
		"provider_id": req.ProviderId,
		"service_id":  req.ServiceId,
		"status":      req.GetStatus(),
		"scheduled_time": bson.M{
			"start_time": req.GetScheduledTime().GetStartTime(),
			"end_time":   req.GetScheduledTime().GetEndTime(),
		},
		"total_price": req.GetTotalPrice(),
		"location": bson.M{
			"city":    req.GetLocation().GetCity(),
			"country": req.GetLocation().GetCountry(),
		},
	}

	resp, err := collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	oid, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert objectid to hex: %v", resp.InsertedID)
	}

	return &pb.CreateBookingResponse{
		Id:         oid.Hex(),
		UserId:     req.UserId,
		ProviderId: req.ProviderId,
		ServiceId:  req.ServiceId,
		Status:     req.GetStatus(),
		TotalPrice: req.TotalPrice,
		Location: &pb.Location{
			City:    req.GetLocation().GetCity(),
			Country: req.GetLocation().GetCountry(),
		},
	}, nil
}

func (b *BookingRepo) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var booking struct {
		UserId        string `bson:"user_id"`
		ProviderId    string `bson:"provider_id"`
		ServiceId     string `bson:"service_id"`
		Status        string `bson:"status"`
		ScheduledTime struct {
			StartTime string `bson:"start_time"`
			EndTime   string `bson:"end_time"`
		} `bson:"scheduled_time"`
		TotalPrice float64 `bson:"total_price"`
		Location   struct {
			City    string `bson:"city"`
			Country string `bson:"country"`
		} `bson:"location"`
		Id primitive.ObjectID `bson:"_id"`
	}

	err = collection.FindOne(ctx, filter).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("booking not found")
		}
		return nil, err
	}

	resp := &pb.GetBookingResponse{
		UserId:     booking.UserId,
		ProviderId: booking.ProviderId,
		ServiceId:  booking.ServiceId,
		Status:     booking.Status,
		ScheduledTime: &pb.ScheduledTime{
			StartTime: booking.ScheduledTime.StartTime,
			EndTime:   booking.ScheduledTime.EndTime,
		},
		TotalPrice: float32(booking.TotalPrice),
		Location: &pb.Location{
			City:    booking.Location.City,
			Country: booking.Location.Country,
		},
		Id: booking.Id.Hex(),
	}

	return resp, nil
}

func (b *BookingRepo) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": oid,
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":     req.GetUserId(),
			"provider_id": req.GetProviderId(),
			"service_id":  req.GetServiceId(),
			"status":      req.GetStatus(),
			"total_price": req.GetTotalPrice(),
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update booking: %v", err)
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("booking not found")
	}

	var updatedBooking struct {
		UserId     string             `bson:"user_id"`
		ProviderId string             `bson:"provider_id"`
		ServiceId  string             `bson:"service_id"`
		Status     string             `bson:"status"`
		TotalPrice float64            `bson:"total_price"`
		Id         primitive.ObjectID `bson:"_id"`
		UpdatedAt  string             `bson:"updated_at"`
	}

	err = collection.FindOne(ctx, filter).Decode(&updatedBooking)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated booking: %v", err)
	}

	return &pb.UpdateBookingResponse{
		UserId:     updatedBooking.UserId,
		ProviderId: updatedBooking.ProviderId,
		ServiceId:  updatedBooking.ServiceId,
		Status:     updatedBooking.Status,
		TotalPrice: float32(updatedBooking.TotalPrice),
		UpdatedAt:  updatedBooking.UpdatedAt,
	}, nil
}

func (b *BookingRepo) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	collection := b.db.Collection("booking")

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, err
	}

	return &pb.CancelBookingResponse{
		Message: "SUCCESS",
	}, nil
}

func (b *BookingRepo) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	collection := b.db.Collection("booking")

	// MongoDB qidiruv parametrlari
	findOptions := options.Find()
	findOptions.SetLimit(int64(req.GetLimit()))
	findOptions.SetSkip(int64(req.GetOffset()))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to list bookings: %v", err)
	}
	defer cursor.Close(ctx)

	var bookings []*pb.GetBookingResponse

	for cursor.Next(ctx) {
		var booking struct {
			UserId        string `bson:"user_id"`
			ProviderId    string `bson:"provider_id"`
			ServiceId     string `bson:"service_id"`
			Status        string `bson:"status"`
			ScheduledTime struct {
				StartTime string `bson:"start_time"`
				EndTime   string `bson:"end_time"`
			} `bson:"scheduled_time"`

			TotalPrice float64 `bson:"total_price"`
			Location   struct {
				City    string `bson:"city"`
				Country string `bson:"country"`
			} `bson:"location"`
			Id primitive.ObjectID `bson:"_id"`
		}

		err := cursor.Decode(&booking)
		if err != nil {
			return nil, fmt.Errorf("failed to decode booking: %v", err)
		}

		// Protobuf Booking message'iga o'tkazish
		bookings = append(bookings, &pb.GetBookingResponse{
			UserId:     booking.UserId,
			ProviderId: booking.ProviderId,
			ServiceId:  booking.ServiceId,
			Status:     booking.Status,
			ScheduledTime: &pb.ScheduledTime{
				StartTime: booking.ScheduledTime.StartTime,
				EndTime:   booking.ScheduledTime.EndTime,
			},
			TotalPrice: float32(booking.TotalPrice),
			Location: &pb.Location{
				City:    booking.Location.City,
				Country: booking.Location.Country,
			},
			Id: booking.Id.Hex(),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.ListBookingsResponse{
		Bookings: bookings,
	}, nil
}
