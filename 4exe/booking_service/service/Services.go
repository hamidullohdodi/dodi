package service

import (
	pb "booking_service/genproto/booking"
)

type MainService interface {
	BookingService() pb.BookingServer
}

type mainService struct {
	booking pb.BookingServer
}

func (m *mainService) BookingService() pb.BookingServer {
	return m.booking
}
