package models

type DoctorSignUp struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password"`
	ConfirmPassword   string `json:"confirm_password"`
	Specialization    string `json:"specialization"`
	YearsOfExperience int32  `json:"years_of_experience"`
	LicenseNumber     string `json:"license_number"`
	Fees              int64  `json:"fees"`
}
type DoctorDetail struct {
	Id                uint   `json:"id"`
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Specialization    string `json:"specialization"`
	YearsOfExperience int32  `json:"years_of_experience"`
	LicenseNumber     string `json:"license_number"`
	Fees              int64  `json:"fees"`
}
type DoctorSignUpResponse struct {
	DoctorDetail DoctorDetail
	AccessToken  string
	RefreshToken string
}
type DoctorLogin struct {
	Email    string
	Password string
}
type DoctorDetails struct {
	Id                uint
	FullName          string
	Email             string
	PhoneNumber       string
	Password          string
	Specialization    string
	YearsOfExperience int32
	LicenseNumber     string
	Fees              int64
}
type DoctorsDetails struct {
	DoctorDetail DoctorDetail
	Rating       int32
}
type UpdateDoctor struct {
	FullName          string `json:"full_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Specialization    string `json:"specialization"`
	YearsOfExperience int32  `json:"years_of_experience"`
	Fees              int64
}
type BookingDoctorDetails struct {
	Doctorid    uint
	DoctorName  string
	DoctorEmail string
	Fees        uint64
}
type DoctorPaymentDetail struct {
	Doctor_id  int
	DoctorName string
	Fees       uint64
}
