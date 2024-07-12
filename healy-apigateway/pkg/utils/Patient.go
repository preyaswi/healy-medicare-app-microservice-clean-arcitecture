package models

type GoogleUserInfo struct {
	ID           string `json:"id"` // Google's unique user ID
	Email        string `json:"email"`
	Name         string `json:"name"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
	TokenExpiry  string `json:"tokenexpiry"`
}
type GoogleSignupdetailResponse struct {
	Id       string
	GoogleId string
	FullName string
	Email    string
}

type SignupdetailResponse struct {
	Id            string
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type TokenPatient struct {
	Patient      GoogleSignupdetailResponse
	AccessToken  string
	RefreshToken string
}

// PatientDetails represents the structure for patient details.
// swagger:model PatientDetails
type PatientDetails struct {
    // The full name of the patient
    // example: John Doe
    Fullname string `json:"fullname"`
    
    // The email address of the patient
    // example: john.doe@example.com
    Email string `json:"email"`
    
    // The gender of the patient
    // example: male
    Gender string `json:"gender"`
    
    // The contact number of the patient
    // example: +1234567890
    Contactnumber string `json:"contactnumber"`
}
type Patient struct {
	BookingId     uint
	SlotId        uint
	PaymentStatus string
	PatientId     uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
