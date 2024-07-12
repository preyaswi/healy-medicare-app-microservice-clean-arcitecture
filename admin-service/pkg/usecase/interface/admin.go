package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
	"time"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
	AddToBooking(patientid string, doctorid int) error
	CancelBooking(patientid string, bookingid int) error
	MakePaymentRazorpay(bookingid int) (domain.Booking, string, error)
	VerifyPayment(booking_id int) error
	GetPaidPatients(doctor_id int) ([]models.BookedPatient, error)
	CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error)

	SetDoctorAvailability(availabiity models.SetAvailability) (string, error)
	GetDoctorAvailability(dotctorid int, date time.Time) ([]models.AvailableSlots, error)

	BookSlot(patientid string, bookingid, slotid int) error
	BookDoctor(patientid string, slotid int) (domain.Booking, string, error)
	VerifyandCalenderCreation(bookingid int, paymentid, razorid string) error
}
