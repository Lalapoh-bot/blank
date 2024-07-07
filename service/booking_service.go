package service

import (
	"booking-system/model"
	"booking-system/repository"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{bookingRepo}
}

func (s *bookingService) CreateBooking(booking *model.Booking) error {
	booking.ID = uuid.New().String()
	booking.StatusPembayaran = "partially_paid"
	booking.PembayaranSebagian = booking.TotalPembayaran * 0.5
	return s.bookingRepo.CreateBooking(booking)
}
