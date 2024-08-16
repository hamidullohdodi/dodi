package repo

import (
	"context"

	"booking_service/config"
	pb "booking_service/genproto/booking"
	mongodb "booking_service/storage/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Booking() StorageInterfase
}
type StorageInterfase interface {
	CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error)
	GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error)
	UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error)
	CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error)
	ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error)

	CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error)
	UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error)
	DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error)
	ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error)
	SearchServices(ctx context.Context, req *pb.SearchServicesRequest) (*pb.SearchServicesResponse, error)

	CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error)
	GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error)
	ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error)
	UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error)
	DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error)

	CreateProvider(ctx context.Context, req *pb.CreateProviderRequest) (*pb.CreateProviderResponse, error)
	GetProvider(ctx context.Context, req *pb.GetProviderRequest) (*pb.GetProviderResponse, error)
	UpdateProvider(ctx context.Context, req *pb.UpdateProviderRequest) (*pb.UpdateProviderResponse, error)
	DeleteProvider(ctx context.Context, req *pb.DeleteProviderRequest) (*pb.DeleteProviderResponse, error)
	ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error)
	SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error)

	CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error)
	UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error)
	DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error)
	ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error)
}

type Store struct {
	Db *mongo.Database
}

func (s *Store) Booking() StorageInterfase {
	return mongodb.NewBookingRepo(s.Db)
}

func NewStorage() IStorage {
	cfg := config.Load()
	db, err := mongodb.NewMongoClient(cfg)
	if err != nil {
		return nil
	}
	return &Store{Db: db}
}
