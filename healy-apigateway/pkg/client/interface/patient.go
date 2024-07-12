package interfaces

import (
	models "healy-apigateway/pkg/utils"
)

type PatientClient interface {
	GoogleSignIn(googleID, email, name, accesstoken, refreshtoken, tokenexpiry string) (models.TokenPatient, error)
	PatientDetails(user_id string) (models.SignupdetailResponse, error)
	UpdatePatientDetails(patient models.PatientDetails, patient_id string) (models.PatientDetails, error)
	ListPatients() ([]models.SignupdetailResponse, error)
}
