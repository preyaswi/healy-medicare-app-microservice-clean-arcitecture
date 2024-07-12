package interfaces

import models "healy-apigateway/pkg/utils"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)

	GetPaidPatients(doctor_id int) ([]models.Patient, error)
	CreatePrescription(prescription models.PrescriptionRequest) (models.CreatedPrescription, error)

	SetDoctorAvailability(setreq models.SetAvailability, doctorId int) (string, error)
	GetDoctorAvailability(doctorid int, date string) ([]models.GetAvailability, error)

	BookDoctor(patientid string, slotid int) (models.CombinedBookingDetails, string, error)
	VerifyandCalenderCreation(bookingId int, paymentid, razorid string) error
}
