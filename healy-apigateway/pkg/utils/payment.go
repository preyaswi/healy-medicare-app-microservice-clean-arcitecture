package models

type CombinedBookingDetails struct {
	BookingId     uint
	PatientName     string
	DoctorId      uint
	DoctorName    string
	DoctorEmail   string
	Fees          uint64
	PaymentStatus string
}
