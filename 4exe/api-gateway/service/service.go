package service

import (
	"api-geteway/internal/config"
	booking "api-geteway/genproto/booking"
	user "api-geteway/genproto/user"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceManager interface defines the methods to access the Booking and User services.
type ServiceManager interface {
	BookingService() booking.BookingClient
	UserService() user.AuthClient
}

// Service struct implements the ServiceManager interface.
type Service struct {
	config         *config.Config
	bookingService booking.BookingClient
	userService    user.AuthClient
}

// NewService creates a new Service with initialized gRPC clients for Booking and User services.
func NewService(cfg *config.Config) (*Service, error) {
	bookingConn, err := grpc.NewClient("booking:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userConn, err := grpc.NewClient("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return &Service{
		config:         cfg,
		bookingService: booking.NewBookingClient(bookingConn),
		userService:    user.NewAuthClient(userConn),
	}, nil
}

// BookingService returns the BookingServiceClient.
func (s *Service) BookingService() booking.BookingClient {
	return s.bookingService
}

// UserService returns the UserServiceClient.
func (s *Service) UserService() user.AuthClient {
	return s.userService
}

// Booking-related methods

func (s *Service) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {
	return s.BookingService().CreateBooking(ctx, req)
}

func (s *Service) GetBooking(ctx context.Context, req *booking.GetBookingRequest) (*booking.GetBookingResponse, error) {
	return s.BookingService().GetBooking(ctx, req)
}

func (s *Service) UpdateBooking(ctx context.Context, req *booking.UpdateBookingRequest) (*booking.UpdateBookingResponse, error) {
	return s.BookingService().UpdateBooking(ctx, req)
}

func (s *Service) CancelBooking(ctx context.Context, req *booking.CancelBookingRequest) (*booking.CancelBookingResponse, error) {
	return s.BookingService().CancelBooking(ctx, req)
}

func (s *Service) ListBookings(ctx context.Context, req *booking.ListBookingsRequest) (*booking.ListBookingsResponse, error) {
	return s.BookingService().ListBookings(ctx, req)
}

// Service-related methods

func (s *Service) CreateService(ctx context.Context, req *booking.CreateServiceRequest) (*booking.CreateServiceResponse, error) {
	return s.BookingService().CreateService(ctx, req)
}

func (s *Service) SearchService(ctx context.Context, req *booking.SearchServicesRequest) (*booking.SearchServicesResponse, error) {
	return s.BookingService().SearchServices(ctx, req)
}

func (s *Service) UpdateService(ctx context.Context, req *booking.UpdateServiceRequest) (*booking.UpdateServiceResponse, error) {
	return s.BookingService().UpdateService(ctx, req)
}

func (s *Service) ListServices(ctx context.Context, req *booking.ListServicesRequest) (*booking.ListServicesResponse, error) {
	return s.BookingService().ListServices(ctx, req)
}

func (s *Service) DeleteService(ctx context.Context, req *booking.DeleteServiceRequest) (*booking.DeleteServiceResponse, error) {
	return s.BookingService().DeleteService(ctx, req)
}

// Provider-related methods

func (s *Service) CreateProvider(ctx context.Context, req *booking.CreateProviderRequest) (*booking.CreateProviderResponse, error) {
	return s.BookingService().CreateProvider(ctx, req)
}

func (s *Service) GetProvider(ctx context.Context, req *booking.GetProviderRequest) (*booking.GetProviderResponse, error) {
	return s.BookingService().GetProvider(ctx, req)
}

func (s *Service) ListProviders(ctx context.Context, req *booking.ListProvidersRequest) (*booking.ListProvidersResponse, error) {
	return s.BookingService().ListProviders(ctx, req)
}

func (s *Service) UpdateProvider(ctx context.Context, req *booking.UpdateProviderRequest) (*booking.UpdateProviderResponse, error) {
	return s.BookingService().UpdateProvider(ctx, req)
}

func (s *Service) DeleteProvider(ctx context.Context, req *booking.DeleteProviderRequest) (*booking.DeleteProviderResponse, error) {
	return s.BookingService().DeleteProvider(ctx, req)
}

func (s *Service) SearchProviders(ctx context.Context, req *booking.SearchProvidersRequest) (*booking.SearchProvidersResponse, error) {
	return s.BookingService().SearchProviders(ctx, req)
}

// Review-related methods

func (s *Service) CreateReview(ctx context.Context, req *booking.CreateReviewRequest) (*booking.CreateReviewResponse, error) {
	return s.BookingService().CreateReview(ctx, req)
}

func (s *Service) ListReviews(ctx context.Context, req *booking.ListReviewsRequest) (*booking.ListReviewsResponse, error) {
	return s.BookingService().ListReviews(ctx, req)
}

func (s *Service) UpdateReview(ctx context.Context, req *booking.UpdateReviewRequest) (*booking.UpdateReviewResponse, error) {
	return s.BookingService().UpdateReview(ctx, req)
}

func (s *Service) DeleteReview(ctx context.Context, req *booking.DeleteReviewRequest) (*booking.DeleteReviewResponse, error) {
	return s.BookingService().DeleteReview(ctx, req)
}

// Payment-related methods

func (s *Service) CreatePayment(ctx context.Context, req *booking.CreatePaymentRequest) (*booking.CreatePaymentResponse, error) {
	return s.BookingService().CreatePayment(ctx, req)
}

func (s *Service) GetPayment(ctx context.Context, req *booking.GetPaymentRequest) (*booking.GetPaymentResponse, error) {
	return s.BookingService().GetPayment(ctx, req)
}

func (s *Service) ListPayments(ctx context.Context, req *booking.ListPaymentsRequest) (*booking.ListPaymentsResponse, error) {
	return s.BookingService().ListPayments(ctx, req)
}

func (s *Service) UpdatePayment(ctx context.Context, req *booking.UpdatePaymentRequest) (*booking.UpdatePaymentResponse, error) {
	return s.BookingService().UpdatePayment(ctx, req)
}

func (s *Service) DeletePayment(ctx context.Context, req *booking.DeletePaymentRequest) (*booking.DeletePaymentResponse, error) {
	return s.BookingService().DeletePayment(ctx, req)
}
