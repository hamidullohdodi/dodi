package mongodb

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *BookingRepo) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	collection := b.db.Collection("reviews")

	// Yozuvni yaratish
	review := bson.M{
		"user_id":    req.GetUserId(),
		"booking_id": req.GetBookingId(),
		"rating":     req.GetRating(),
		"comment":    req.GetComment(),
	}

	// Yozuvni ma'lumotlar bazasiga saqlash
	resp, err := collection.InsertOne(ctx, review)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %v", err)
	}
	// ID'ni olish va formatlash
	_, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID: %v", resp.InsertedID)
	}
	// Javobni tayyorlash
	return &pb.CreateReviewResponse{
		UserId:    req.GetUserId(),
		BookingId: req.GetBookingId(),
		Rating:    req.GetRating(),
		Comment:   req.GetComment(),
	}, nil
}



func (b *BookingRepo) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	collection := b.db.Collection("reviews")
	// MongoDB qidiruv parametrlari
	filter := bson.M{}
	if req.Limit != 0 {
		filter["limit"] = req.Limit
	}
	if req.Offset != 0 {
		filter["offset"] = req.Offset
	}

	// MongoDB qidiruv
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list reviews: %v", err)
	}
	defer cursor.Close(ctx)
	var reviews []*pb.Review
	for cursor.Next(ctx) {
		var review pb.Review
		err := cursor.Decode(&review)
		if err != nil {
			return nil, fmt.Errorf("failed to decode review: %v", err)
		}
		reviews = append(reviews, &review)
	}
	return &pb.ListReviewsResponse{Reviews: reviews}, nil
}

func (b *BookingRepo) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	collection := b.db.Collection("reviews")
	// ID'ni olish va formatlash
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}
	// Yozuvni yangilash
	update := bson.M{
		"$set": bson.M{
			"user_id":     req.GetUserId(),
			"provider_id": req.GetProviderId(),
			"rating":      req.GetRating(),
			"comment":     req.GetComment(),
		},
	}

	// Yozuvni yangilash
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update review: %v", err)
	}
	// Javobni tayyorlash
	return &pb.UpdateReviewResponse{
		UserId:     req.GetUserId(),
		ProviderId: req.GetProviderId(),
		Rating:     req.GetRating(),
		Comment:    req.GetComment(),
	}, nil

}

func (b *BookingRepo) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	collection := b.db.Collection("reviews")
	// ID'ni olish va formatlash
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}
	// Yozuvni o'chirish
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, fmt.Errorf("failed to delete review: %v", err)
	}
	// Javobni tayyorlash
	return &pb.DeleteReviewResponse{}, nil
}
