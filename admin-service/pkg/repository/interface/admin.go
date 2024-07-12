package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
	"time"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)

	AddToBooking(patientid string, doctordetail models.BookingDoctorDetails) error
	AddDetailsToBooking(patientid string, doctordetail models.BookingDoctorDetails, slotid int) (int, error)
	GetBookingByID(bookingid int) (domain.Booking, error)
	UpdateSlotAvailability(slotid int) error
	GetAvailabilityByID(slotID int) (domain.Availability, error)
	StoreEventDetails(event domain.Event) error
	RemoveBooking(bookingID int) error

	AddRazorPayDetails(bookingid uint, razorPaypaymentID string) error
	CheckPaymentStatus(bookingid int) (string, error)
	UpdatePaymentStatus(bookingid int, status string) error
	GetPaidBookingsByDoctorID(doctorId int) ([]domain.Booking, error)
	CheckPatientPayment(patientid string, bookingid int) (bool, error)
	CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error)

	SetDoctorAvailability(availabiity models.SetAvailability) (string, error)
	GetDoctorAvailability(doctor_id int, date time.Time) ([]models.AvailableSlots, error)
	CheckSlotAvailability(slotid int) (bool, error)
	BookSlot(bookingid, slotid int) error
	MarkSlotAsBooked(slotid int) error

	GetDoctorIdFromSlotId(slotid int) (int, error)
	UpdatePaymentDetails(bookingid int, paymentid string) error
}
