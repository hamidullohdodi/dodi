package service

import (
	pb "booking_service/genproto/booking"
	"context"
	"fmt"

	Storage "booking_service/storage/repo"
)

type BookingService struct {
	pb.UnimplementedBookingServer
	storage Storage.IStorage
}

func NewBookingService(storage Storage.IStorage) *BookingService {
	return &BookingService{
		storage: storage,
	}
}

func (b *BookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	booking, err := b.storage.Booking().CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	booking, err := b.storage.Booking().GetBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	booking, err := b.storage.Booking().UpdateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	booking, err := b.storage.Booking().CancelBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	bookings, err := b.storage.Booking().ListBookings(ctx, req)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}



func (b *BookingService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	booking, err := b.storage.Booking().CreateReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	booking, err := b.storage.Booking().UpdateReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	booking, err := b.storage.Booking().ListReviews(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	booking, err := b.storage.Booking().DeleteReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}




func (b *BookingService) CreateProvider(ctx context.Context, req *pb.CreateProviderRequest) (*pb.CreateProviderResponse, error) {
	booking, err := b.storage.Booking().CreateProvider(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) UpdateProvider(ctx context.Context, req *pb.UpdateProviderRequest) (*pb.UpdateProviderResponse, error) {
	booking, err := b.storage.Booking().UpdateProvider(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) GetProvider(ctx context.Context, req *pb.GetProviderRequest) (*pb.GetProviderResponse, error) {
	booking, err := b.storage.Booking().GetProvider(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) ListProviders(ctx context.Context, req *pb.ListProvidersRequest) (*pb.ListProvidersResponse, error) {
	booking, err := b.storage.Booking().ListProviders(ctx, req)
	fmt.Println(111)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) DeleteProvider(ctx context.Context, req *pb.DeleteProviderRequest) (*pb.DeleteProviderResponse, error) {
	booking, err := b.storage.Booking().DeleteProvider(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) SearchProviders(ctx context.Context, req *pb.SearchProvidersRequest) (*pb.SearchProvidersResponse, error) {
	booking, err := b.storage.Booking().SearchProviders(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}


func (b *BookingService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	booking, err := b.storage.Booking().CreatePayment(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error) {
	booking, err := b.storage.Booking().UpdatePayment(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	booking, err := b.storage.Booking().GetPayment(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
	booking, err := b.storage.Booking().ListPayments(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	booking, err := b.storage.Booking().DeletePayment(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}


func (b *BookingService) CreateService(ctx context.Context, req *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	booking, err := b.storage.Booking().CreateService(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) UpdateService(ctx context.Context, req *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error) {
	booking, err := b.storage.Booking().UpdateService(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) DeleteService(ctx context.Context, req *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error) {
	booking, err := b.storage.Booking().DeleteService(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	booking, err := b.storage.Booking().ListServices(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) SearchServices(ctx context.Context, req *pb.SearchServicesRequest) (*pb.SearchServicesResponse, error) {
	booking, err := b.storage.Booking().SearchServices(ctx, req)
	if err != nil {
		return nil, err
	}
	return booking, nil
}