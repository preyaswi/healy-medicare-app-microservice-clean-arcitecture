package models

import "time"

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
type AdminSignUp struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type BookingDoctorDetails struct {
	Doctorid    uint
	DoctorName  string
	DoctorEmail string
	Fees        uint64
}
type Paymentreq struct {
	Bookingid uint
}
type PaymentDetails struct {
	PatientId     uint
	DoctorId      uint
	Fees          uint64
	PaymentStatus string
}

type SetAvailability struct {
	DoctorId  int
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
}
type AvailableSlots struct {
	Slot_id  int
	Time     string
	IsBooked bool
}
