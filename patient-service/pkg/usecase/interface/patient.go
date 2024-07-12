package usecaseint

import (
	"patient-service/pkg/models"
)

type PatientUseCase interface {
	GoogleSignIn(googleid, googlename, googleEmail, googleaccesstoken, googlerefreshtoken, googletokenexpiry string) (models.TokenPatient, error)
	IndPatientDetails(patient_id string) (models.SignupdetailResponse, error)
	UpdatePatientDetails(patient models.SignupdetailResponse) (models.SignupdetailResponse, error)

	ListPatients() ([]models.SignupdetailResponse, error)

	GetPatientGoogleDetailsByID(patientid string) (models.GooglePatientDetails, error)
	UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error
}
