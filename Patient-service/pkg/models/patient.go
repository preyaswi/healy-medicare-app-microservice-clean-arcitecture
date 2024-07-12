package models

type PatientSignUp struct {
	Fullname        string
	Email           string
	Password        string
	Confirmpassword string
	Gender          string
	Contactnumber   string
}
type GoogleSignupdetailResponse struct {
	Id       string
	GoogleId string
	FullName string
	Email    string
}
type SignupdetailResponse struct {
	Id            string `json:"id"`
	Fullname      string `json:"fullname"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Contactnumber string `json:"contactnumber"`
}

type TokenPatient struct {
	Patient      GoogleSignupdetailResponse
	AccessToken  string
	RefreshToken string
}

type PatientLogin struct {
	Email    string
	Password string
}
type PatientDetails struct {
	Id            uint
	Fullname      string
	Email         string
	Password      string
	Gender        string
	Contactnumber string
}
type GooglePatientDetails struct {
	GoogleID     string
	GoogleEmail  string
	AccessToken  string
	RefreshToken string
	TokenExpiry  string
}
