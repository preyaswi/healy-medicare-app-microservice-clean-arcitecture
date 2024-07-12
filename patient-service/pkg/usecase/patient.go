package usecase

import (
	"errors"
	"fmt"
	"patient-service/pkg/helper"
	"patient-service/pkg/models"
	interfaces "patient-service/pkg/repository/interface"
	usecaseint "patient-service/pkg/usecase/interface"
)

type patientUseCase struct {
	patientRepository interfaces.PatientRepository
}

func NewPatientUseCase(repository interfaces.PatientRepository) usecaseint.PatientUseCase {
	return &patientUseCase{
		patientRepository: repository,
	}
}

func (pr *patientUseCase) GoogleSignIn(googleid, googlename, googleEmail, googleaccesstoken, googlerefreshtoken, googletokenexpiry string) (models.TokenPatient, error) {
	patient, err := pr.patientRepository.FindOrCreatePatientByGoogleID(googleid, googleEmail, googlename, googleaccesstoken, googlerefreshtoken, googletokenexpiry)
	if err != nil {
		return models.TokenPatient{}, err
	}
	accesstoken, err := helper.GenerateAccessToken(patient)
	if err != nil {
		return models.TokenPatient{}, err
	}
	refreshToken, err := helper.GenerateRefreshToken(patient)
	if err != nil {
		return models.TokenPatient{}, err
	}
	return models.TokenPatient{
		Patient:      patient,
		AccessToken:  accesstoken,
		RefreshToken: refreshToken,
	}, nil

}
func (pr *patientUseCase) IndPatientDetails(patient_id string) (models.SignupdetailResponse, error) {
	pateint, err := pr.patientRepository.IndPatientDetails(patient_id)
	if err != nil {
		return models.SignupdetailResponse{}, nil
	}
	return pateint, nil
}
func (pr *patientUseCase) UpdatePatientDetails(patient models.SignupdetailResponse) (models.SignupdetailResponse, error) {
	fmt.Println(patient, "patient details")

	// Update email if provided and not already existing
	if patient.Email != "" {
		userExist := pr.patientRepository.CheckPatientAvailability(patient.Email)
		if userExist {
			return models.SignupdetailResponse{}, errors.New("user already exists, choose a different email")
		}
		if err := pr.patientRepository.UpdatePatientEmail(patient.Email, patient.Id); err != nil {
			return models.SignupdetailResponse{}, err
		}
	}

	// Update contact number if provided and not already existing
	if patient.Contactnumber != "" {
		userExistByPhone, err := pr.patientRepository.CheckPatientExistsByPhone(patient.Contactnumber)
		fmt.Println(userExistByPhone, "userExistByPhone")
		if err != nil {
			return models.SignupdetailResponse{}, errors.New("error with server")
		}
		if userExistByPhone != nil {
			return models.SignupdetailResponse{}, errors.New("user with this phone number already exists")
		}
		if err := pr.patientRepository.UpdatePatientPhone(patient.Contactnumber, patient.Id); err != nil {
			return models.SignupdetailResponse{}, err
		}
	}

	// Update name if provided
	if patient.Fullname != "" {
		if err := pr.patientRepository.UpdateName(patient.Fullname, patient.Id); err != nil {
			return models.SignupdetailResponse{}, err
		}
	}
	if patient.Gender != "" {

		if err := pr.patientRepository.UpdateGender(patient.Gender, patient.Id); err != nil {
			return models.SignupdetailResponse{}, err

		}
	}
	// Retrieve and return updated user details
	updatedPatient, err := pr.patientRepository.IndPatientDetails(patient.Id)
	if err != nil {
		return models.SignupdetailResponse{}, errors.New("could not get user details")
	}

	return updatedPatient, nil
}
func (pr *patientUseCase) ListPatients() ([]models.SignupdetailResponse, error) {
	patients, err := pr.patientRepository.ListPatients()
	if err != nil {
		return []models.SignupdetailResponse{}, err
	}
	return patients, nil
}
func (pr *patientUseCase) GetPatientGoogleDetailsByID(patientid string) (models.GooglePatientDetails, error) {
	patientdetails, err := pr.patientRepository.GetPatientGoogleDetailsByID(patientid)
	if err != nil {
		return models.GooglePatientDetails{}, err
	}
	return patientdetails, nil
}
func (pr *patientUseCase) UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error {
	err := pr.patientRepository.UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry)
	if err != nil {
		return err
	}
	return nil
}
