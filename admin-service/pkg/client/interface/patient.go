package interfaces

import "healy-admin/pkg/utils/models"

type NewPatientClient interface {
	GetPatientByID(patientid string) (models.Patient, error)
	GetPatientGoogleDetailsByID(patientid string) (models.GooglePatientDetails, error)
	UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error
}
