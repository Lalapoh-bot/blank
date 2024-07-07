package model

import "time"

type Booking struct {
	ID                 string    `json:"id"`
	CustomerID         string    `json:"customer_id"`
	EmployeeID         string    `json:"employee_id"`
	LapanganID         string    `json:"lapangan_id"`
	BookingDate        time.Time `json:"booking_date"`
	Jam                string    `json:"jam"`
	TotalPembayaran    float64   `json:"total_pembayaran"`
	PembayaranSebagian float64   `json:"pembayaran_sebagian"`
	StatusPembayaran   string    `json:"status_pembayaran"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
