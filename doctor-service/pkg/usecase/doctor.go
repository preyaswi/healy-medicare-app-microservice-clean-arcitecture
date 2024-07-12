package usecase

import (
	"doctor-service/pkg/helper"
	"doctor-service/pkg/models"
	interfaces "doctor-service/pkg/repository/interface"
	usecaseint "doctor-service/pkg/usecase/interface"
	"errors"
	"strconv"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type doctorUseCase struct {
	doctorRepository interfaces.DoctorRepository
}

func NewDoctorUseCase(repository interfaces.DoctorRepository) usecaseint.DoctorUseCase {
	return &doctorUseCase{
		doctorRepository: repository,
	}
}
func (du *doctorUseCase) DoctorSignUp(doctor models.DoctorSignUp) (models.DoctorSignUpResponse, error) {
	email, err := du.doctorRepository.CheckDoctorExistsByEmail(doctor.Email)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if email != nil {
		return models.DoctorSignUpResponse{}, errors.New("user with this email is already exists")
	}

	phone, err := du.doctorRepository.CheckDoctorExistsByPhone(doctor.PhoneNumber)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if phone != nil {
		return models.DoctorSignUpResponse{}, errors.New("user with this phone is already exists")
	}
	if doctor.Password != doctor.ConfirmPassword {
		return models.DoctorSignUpResponse{}, errors.New("confirm password is not matching")
	}

	hashPassword, err := helper.PasswordHash(doctor.Password)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error in hashing password")
	}
	doctor.Password = hashPassword
	data, err := du.doctorRepository.CheckDoctorexistsByLicenseNumber(doctor.LicenseNumber)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with the server")
	}
	if data != nil {
		return models.DoctorSignUpResponse{}, errors.New("doctor with this license number is already exists,please double check")
	}
	DoctorData, err := du.doctorRepository.DoctorSignup(doctor)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("could not add the doctor")
	}
	accessToken, err := helper.GenerateAccessToken(DoctorData)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(DoctorData)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create refresh token due to error")
	}
	return models.DoctorSignUpResponse{
		DoctorDetail: DoctorData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (du *doctorUseCase) DoctorLogin(doctor models.DoctorLogin) (models.DoctorSignUpResponse, error) {
	data, err := du.doctorRepository.CheckDoctorExistsByEmail(doctor.Email)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("error with server")
	}
	if data == nil {
		return models.DoctorSignUpResponse{}, errors.New("email doesn't exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(doctor.Password))

	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("password not matching")
	}
	var doctorDetail models.DoctorDetail
	err = copier.Copy(&doctorDetail, &data)
	if err != nil {
		return models.DoctorSignUpResponse{}, err
	}
	accessToken, err := helper.GenerateAccessToken(doctorDetail)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(doctorDetail)
	if err != nil {
		return models.DoctorSignUpResponse{}, errors.New("counldn't create refreshtoken due to internal error")
	}
	return models.DoctorSignUpResponse{
		DoctorDetail: doctorDetail,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (du *doctorUseCase) DoctorsList() ([]models.DoctorsDetails, error) {
	doctors, err := du.doctorRepository.GetDoctorsDetail()
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
func (du *doctorUseCase) IndividualDoctor(doctor_id string) (models.DoctorsDetails, error) {
	doctor, err := du.doctorRepository.ShowIndividualDoctor(doctor_id)
	if err != nil {
		return models.DoctorsDetails{}, err
	}
	return doctor, nil
}
func (du *doctorUseCase) DoctorProfile(id int) (models.DoctorsDetails, error) {
	doctor, err := du.doctorRepository.DoctorProfile(id)
	if err != nil {
		return models.DoctorsDetails{}, err
	}
	return doctor, nil
}
func (du *doctorUseCase) RateDoctor(patientid string, doctorid string, rate uint32) (uint32, error) {
	doctor_id, err := strconv.Atoi(doctorid)
	if err != nil {
		return 0, errors.New("couldn't convert doctor string to int")
	}
	ok, err := du.doctorRepository.CheckDoctorExistbyid(doctor_id)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("doctor with this id is not exsisting")
	}
	rated, err := du.doctorRepository.RateDoctor(patientid, doctorid, rate)
	if err != nil {
		return 0, err
	}
	return uint32(rated), nil

}
func (du *doctorUseCase) UpdateDoctorProfile(doctorID int, doctor models.UpdateDoctor) (models.UpdateDoctor, error) {
	// Check if the email already exists and it's not the same doctor
	if doctor.Email != "" {
		existingDoctor, err := du.doctorRepository.CheckDoctorExistsByEmail(doctor.Email)
		if err != nil {
			return models.UpdateDoctor{}, errors.New("error with server")
		}
		if existingDoctor != nil && existingDoctor.ID != uint(doctorID) {
			return models.UpdateDoctor{}, errors.New("doctor already exists, choose a different email")
		}
	}

	// Check if the phone number already exists and it's not the same doctor
	if doctor.PhoneNumber != "" {
		existingDoctor, err := du.doctorRepository.CheckDoctorExistsByPhone(doctor.PhoneNumber)
		if err != nil {
			return models.UpdateDoctor{}, errors.New("error with server")
		}
		if existingDoctor != nil && existingDoctor.ID != uint(doctorID) {
			return models.UpdateDoctor{}, errors.New("doctor with this phone number already exists")
		}
	}

	// Update fields
	if doctor.Email != "" {
		if err := du.doctorRepository.UpdateDoctorField("email", doctor.Email, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	if doctor.FullName != "" {
		if err := du.doctorRepository.UpdateDoctorField("full_name", doctor.FullName, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	if doctor.PhoneNumber != "" {
		if err := du.doctorRepository.UpdateDoctorField("phone_number", doctor.PhoneNumber, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	if doctor.Specialization != "" {
		if err := du.doctorRepository.UpdateDoctorField("specialization", doctor.Specialization, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	if doctor.YearsOfExperience != 0 {
		if err := du.doctorRepository.UpdateDoctorField("years_of_experience", doctor.YearsOfExperience, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	if doctor.Fees != 0 {
		if err := du.doctorRepository.UpdateDoctorField("fees", doctor.YearsOfExperience, uint(doctorID)); err != nil {
			return models.UpdateDoctor{}, err
		}
	}
	udated, err := du.doctorRepository.DoctorDetails(doctorID)
	if err != nil {
		return models.UpdateDoctor{}, err
	}
	return udated, nil
}
func (du *doctorUseCase) DoctorDetailforBooking(doctorid int) (models.BookingDoctorDetails, error) {
	doctordetail, err := du.doctorRepository.DoctorDetailforBooking(doctorid)
	if err != nil {
		return models.BookingDoctorDetails{}, err
	}
	return doctordetail, nil
}

func (du *doctorUseCase) CheckDoctor(doctorid int) (bool, error) {
	ok, err := du.doctorRepository.CheckDoctorExistbyid(doctorid)
	if err != nil {
		return false, err
	}
	return ok, nil
}
