package mongodb

import (
	"booking_service/genproto/booking"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateBooking(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("create booking", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateCursorResponse(1, "booking.booking", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: mockID},
		}))

		// Initialize the repository
		repo := NewBookingRepo(mt.DB)

		// Create a sample booking request
		req := &booking.CreateBookingRequest{
			UserId:     "user123",
			ProviderId: "provider123",
			ServiceId:  "service123",
			Status:     "scheduled",
			ScheduledTime: &booking.ScheduledTime{
				StartTime: "2024-08-12T10:00:00Z",
				EndTime:   "2024-08-12T11:00:00Z",
			},
			TotalPrice: 100.0,
			Location: &booking.Location{
				City:    "New York",
				Country: "USA",
			},
		}

		// Call the CreateBooking function
		resp, err := repo.CreateBooking(context.Background(), req)

		// Assert no error occurred
		assert.NoError(mt, err)

		// Assert that the response is not nil
		assert.NotNil(mt, resp)

		// Assert the returned ID matches the mock ID
		// assert.Equal(mt, mockID.Hex(), resp.Id)
	})
}

func TestGetBooking(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("get booking", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "booking.booking", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: mockID},
		}))
		// Initialize the repository
		repo := NewBookingRepo(mt.DB)
		// Call the GetBooking function
		resp, err := repo.GetBooking(context.Background(), &booking.GetBookingRequest{Id: mockID.Hex()})
		// Assert no error occurred
		assert.NoError(mt, err)
		// Assert that the response is not nil
		assert.NotNil(mt, resp)
		// Assert the returned ID matches the mock ID
		// assert.Equal(mt, mockID.Hex(), resp.Id)
	})
}

func TestUpdateBooking(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("update booking", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateCursorResponse(1, "booking.booking", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: mockID},
		}))
		// Initialize the repository
		repo := NewBookingRepo(mt.DB)
		// Call the UpdateBooking function
		resp, err := repo.UpdateBooking(context.Background(), &booking.UpdateBookingRequest{Id: mockID.Hex()})
		// Assert no error occurred
		assert.NoError(mt, err)
		// Assert that the response is not nil
		assert.NotNil(mt, resp)
		// Assert the returned ID matches the mock ID
		// assert.Equal(mt, mockID.Hex(), resp.Id)
	})
}

func TestDeleteBooking(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("delete booking", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		// Initialize the repository
		repo := NewBookingRepo(mt.DB)
		// Call the CancelBooking function
		_, err := repo.CancelBooking(context.Background(), &booking.CancelBookingRequest{Id: mockID.Hex()})
		// Assert no error occurred
		assert.NoError(mt, err)
	})
}

func TestListBooking(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("list booking", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "booking.booking", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: mockID},
		}))
		// Initialize the repository
		repo := NewBookingRepo(mt.DB)
		// Call the ListBooking function
		resp, err := repo.ListBookings(context.Background(), &booking.ListBookingsRequest{})
		// Assert no error occurred
		assert.NoError(mt, err)
		// Assert that the response is not nil
		assert.NotNil(mt, resp)
		// Assert the returned ID matches the mock ID
		// assert.Equal(mt, mockID.Hex(), resp.Id)
	})
}

func TestGetBookingByUser(t *testing.T) {
	// Create a new mtest instance
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// Configure the mock to expect an insert and return a mock ID
	mockID := primitive.NewObjectID()
	mt.Run("get booking by user", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "booking.booking", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: mockID},
		}))
		// Initialize the repository
		repo := NewBookingRepo(mt.DB)
		// Call the GetBookingByUser function
		resp, err := repo.GetBooking(context.Background(), &booking.GetBookingRequest{Id: mockID.Hex()})
		// Assert no error occurred
		assert.NoError(mt, err)
		// Assert that the response is not nil
		assert.NotNil(mt, resp)
		// Assert the returned ID matches the mock ID
		// assert.Equal(mt
		// Assert the returned ID matches the mock ID
	})
}
