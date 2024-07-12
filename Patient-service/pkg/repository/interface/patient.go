package interfaces

import (
	"patient-service/pkg/domain"
	"patient-service/pkg/models"
)

type PatientRepository interface {
	FindOrCreatePatientByGoogleID(googleID, email, name, accesstoken, refreshtoken, tokenexpiry string) (models.GoogleSignupdetailResponse, error)
	CheckPatientExistsByEmail(email string) (*domain.Patient, error)
	CheckPatientExistsByPhone(phone string) (*domain.Patient, error)
	FindPatientByEmail(email string) (models.PatientDetails, error)
	IndPatientDetails(patient_id string) (models.SignupdetailResponse, error)
	CheckPatientAvailability(email string) bool
	UpdatePatientEmail(email string, PatientID string) error
	UpdatePatientPhone(phone string, PatientID string) error
	UpdateName(name string, PatientID string) error
	UpdateGender(gender string, patientId string) error
	ListPatients() ([]models.SignupdetailResponse, error)
	GetPatientGoogleDetailsByID(patientid string) (models.GooglePatientDetails, error)
	UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error
}
