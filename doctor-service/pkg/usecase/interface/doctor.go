package usecaseint

import "doctor-service/pkg/models"

type DoctorUseCase interface {
	DoctorSignUp(models.DoctorSignUp) (models.DoctorSignUpResponse, error)
	DoctorLogin(models.DoctorLogin) (models.DoctorSignUpResponse, error)
	DoctorsList() ([]models.DoctorsDetails, error)
	IndividualDoctor(doctor_id string) (models.DoctorsDetails, error)
	DoctorProfile(id int) (models.DoctorsDetails, error)
	RateDoctor(patientid string, doctorid string, rate uint32) (uint32, error)
	UpdateDoctorProfile(doctorid int, body models.UpdateDoctor) (models.UpdateDoctor, error)

	CheckDoctor(doctorid int) (bool, error)
	DoctorDetailforBooking(doctorid int) (models.BookingDoctorDetails, error)
}
