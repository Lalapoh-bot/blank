package repository

import (
	"database/sql"
	"time"

	"booking-system/model"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) error {
	query := `
        INSERT INTO bookings (id, customer_id, employee_id, lapangan_id, booking_date, jam, total_pembayaran, pembayaran_sebagian, status_pembayaran, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `
	_, err := r.db.Exec(query, booking.ID, booking.CustomerID, booking.EmployeeID, booking.LapanganID, booking.BookingDate, booking.Jam, booking.TotalPembayaran, booking.PembayaranSebagian, booking.StatusPembayaran, time.Now(), time.Now())
	return err
}
