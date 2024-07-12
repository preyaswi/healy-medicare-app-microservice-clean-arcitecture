package domain

import (
	"healy-admin/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type Booking struct {
	BookingId     uint   `json:"booking_id" gorm:"primaryKey;not null"`
	PatientId     string `json:"patient_id" gorm:"not null"`
	DoctorId      uint   `json:"doctor_id" gorm:"not null"`
	DoctorName    string `json:"doctor_name" gorm:"not null"`
	DoctorEmail   string `json:"doctor_email" gorm:"not null"`
	Fees          uint64 `json:"fees" gorm:"not null"`
	PaymentStatus string `json:"payment_status" gorm:"default:'not paid'"`
	Slot_id       uint   `json:"slot_id"`
}

type RazerPay struct {
	ID        uint    `json:"id" gorm:"primaryKey;not null"`
	RazorID   string  `json:"razor_id"`
	PaymentID string  `json:"Payment_id"`
	BookingID uint    `json:"Booking_id"`
	Booking   Booking `json:"-" gorm:"foreignKey:BookingID;references:BookingId"`
}
type Prescription struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingID uint   `json:"booking_id" gorm:"not null"`
	DoctorID  uint   `json:"doctor_id" gorm:"not null"`
	PatientID string `json:"patient_id" gorm:"not null"`
	Medicine  string `json:"medicine" gorm:"not null"`
	Dosage    string `json:"dosage" gorm:"not null"`
	Notes     string `json:"notes"`
}

type Availability struct {
	gorm.Model
	DoctorID  uint
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
	IsBooked  bool
}
type Event struct {
	gorm.Model
	BookingID   uint
	PatientID   string
	EventID     string
	Summary     string
	Description string
	Start       time.Time
	End         time.Time
	GuestEmail  string
}
