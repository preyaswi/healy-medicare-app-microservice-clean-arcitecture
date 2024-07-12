package repository

import (
	"encoding/base64"
	"errors"
	"fmt"
	"patient-service/pkg/domain"
	"patient-service/pkg/models"
	interfaces "patient-service/pkg/repository/interface"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type patientRepository struct {
	DB *gorm.DB
}

func NewPatientRepository(DB *gorm.DB) interfaces.PatientRepository {
	return &patientRepository{
		DB: DB,
	}
}
func generateShortUUID() string {
	uuid := uuid.New()
	return base64.RawURLEncoding.EncodeToString(uuid[:])
}
func (ur *patientRepository) FindOrCreatePatientByGoogleID(googleID, email, name, accesstoken, refreshtoken, tokenexpiry string) (models.GoogleSignupdetailResponse, error) {
	var patient domain.Patient
	if err := ur.DB.Where(&domain.Patient{GoogleId: googleID}).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new patient with UUID
			patient = domain.Patient{
				ID:           generateShortUUID(),
				GoogleId:     googleID,
				Email:        email,
				Fullname:     name,
				AccessToken:  accesstoken,
				RefreshToken: refreshtoken,
				TokenExpiry:  tokenexpiry,
			}
			if err := ur.DB.Create(&patient).Error; err != nil {
				return models.GoogleSignupdetailResponse{}, err
			}
		} else {
			return models.GoogleSignupdetailResponse{}, err
		}
	}

	return models.GoogleSignupdetailResponse{
		Id:       patient.ID,
		Email:    patient.Email,
		FullName: patient.Fullname,
		GoogleId: patient.GoogleId,
	}, nil
}

func (ur *patientRepository) CheckPatientExistsByEmail(email string) (*domain.Patient, error) {
	var patient domain.Patient
	res := ur.DB.Where(&domain.Patient{Email: email}).First(&patient)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Patient{}, res.Error
	}
	return &patient, nil
}
func (ur *patientRepository) CheckPatientExistsByPhone(phone string) (*domain.Patient, error) {
	var patient domain.Patient
	res := ur.DB.Where(&domain.Patient{Contactnumber: phone}).First(&patient)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Patient{}, res.Error
	}
	return &patient, nil
}
func (ur *patientRepository) FindPatientByEmail(email string) (models.PatientDetails, error) {
	var patientdetail models.PatientDetails
	err := ur.DB.Raw("SELECT * FROM patient WHERE email=?", email).Scan(&patientdetail).Error
	if err != nil {
		return models.PatientDetails{}, errors.New("error checking user details")
	}
	return patientdetail, nil
}
func (ur *patientRepository) IndPatientDetails(patient_id string) (models.SignupdetailResponse, error) {
	var patient models.SignupdetailResponse
	res := ur.DB.Table("patients").Where("id = ?", patient_id).First(&patient)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return models.SignupdetailResponse{}, errors.New("patient with this id does not exist")
		}
		return models.SignupdetailResponse{}, res.Error
	}

	return patient, nil
}

func (ur *patientRepository) CheckPatientAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from patient where email='%s'", email)
	if err := ur.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func (ur *patientRepository) UpdatePatientEmail(email string, PatientID string) error {

	err := ur.DB.Exec("update patient set email = ? where id = ?", email, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}

func (ur *patientRepository) UpdatePatientPhone(phone string, PatientID string) error {

	err := ur.DB.Exec("update patients set contactnumber = ? where id = ?", phone, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}

func (ur *patientRepository) UpdateName(name string, PatientID string) error {

	err := ur.DB.Exec("update patient set fullname = ? where id = ?", name, PatientID).Error
	if err != nil {
		return err
	}
	return nil

}
func (ur *patientRepository) UpdateGender(gender string, patientid string) error {
	err := ur.DB.Exec("update patients set gender = ? where id = ?", gender, patientid).Error
	if err != nil {
		return err
	}
	return nil
}
func (ur *patientRepository) ListPatients() ([]models.SignupdetailResponse, error) {
	row, err := ur.DB.Raw("select id,fullname,email,gender,contactnumber from patients").Rows()
	if err != nil {
		return []models.SignupdetailResponse{}, err
	}
	defer row.Close()
	var patientDetails []models.SignupdetailResponse
	for row.Next() {
		var patientdetail models.SignupdetailResponse

		// Scan the row into variables
		if err := row.Scan(&patientdetail.Id, &patientdetail.Fullname, &patientdetail.Email, &patientdetail.Gender, &patientdetail.Contactnumber); err != nil {
			return nil, err
		}

		patientDetails = append(patientDetails, patientdetail)
	}
	return patientDetails, nil
}
func (ur *patientRepository) GetPatientGoogleDetailsByID(patientid string) (models.GooglePatientDetails, error) {
	var patient domain.Patient
	var googleDetails models.GooglePatientDetails

	// Find the patient by ID
	err := ur.DB.Where("id = ?", patientid).First(&patient).Error
	if err != nil {
		return googleDetails, err
	}

	// Map the patient details to GooglePatientDetails
	googleDetails = models.GooglePatientDetails{
		GoogleID:     patient.GoogleId,
		GoogleEmail:  patient.Email,
		AccessToken:  patient.AccessToken,
		RefreshToken: patient.RefreshToken,
		TokenExpiry:  patient.TokenExpiry,
	}

	return googleDetails, nil
}
func (ur *patientRepository) UpdatePatientGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error {
	return ur.DB.Model(&domain.Patient{}).Where("google_id = ?", googleID).
		Updates(map[string]interface{}{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"token_expiry":  tokenExpiry,
		}).Error
}
