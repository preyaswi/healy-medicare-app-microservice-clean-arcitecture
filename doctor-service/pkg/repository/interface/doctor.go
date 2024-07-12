package interfaces

import (
	"doctor-service/pkg/domain"
	"doctor-service/pkg/models"
)

type DoctorRepository interface {
	CheckDoctorExistsByEmail(email string) (*domain.Doctor, error)
	CheckDoctorExistsByPhone(phone string) (*domain.Doctor, error)
	CheckDoctorexistsByLicenseNumber(licenceNumber string) (*domain.Doctor, error)
	DoctorSignup(models.DoctorSignUp) (models.DoctorDetail, error)
	GetDoctorsDetail() ([]models.DoctorsDetails, error)
	ShowIndividualDoctor(doctor_id string) (models.DoctorsDetails, error)
	DoctorProfile(id int) (models.DoctorsDetails, error)
	CheckDoctorExistbyid(id int) (bool, error)
	RateDoctor(patient_id string, doctor_id string, rate uint32) (int, error)

	UpdateDoctorField(field string, value interface{}, doctorID uint) error
	DoctorDetails(doctorID int) (models.UpdateDoctor, error)

	DoctorDetailforBooking(doctorid int) (models.BookingDoctorDetails, error)
}
