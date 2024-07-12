package interfaces

import models "healy-apigateway/pkg/utils"

type DoctorClient interface {
	DoctorSignUp(models.DoctorSignUp) (models.DoctorSignUpResponse, error)
	DoctorLogin(models.DoctorLogin) (models.DoctorSignUpResponse, error)
	DoctorsDetails() ([]models.DoctorsDetails, error)
	IndividualDoctor(doctorId string) (models.IndDoctorDetail, error)
	DoctorProfile(id int) (models.IndDoctorDetail, error)
	RateDoctor(patientid string, doctorid string, rate models.Rate) (models.Rate, error)
	UpdateDoctorProfile(doctorid int, body models.DoctorDetails) (models.DoctorDetails, error)
}
