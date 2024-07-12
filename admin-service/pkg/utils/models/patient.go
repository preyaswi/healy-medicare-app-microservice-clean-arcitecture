package models

type Patient struct {
	Id            uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type BookedPatient struct {
	BookingId     int
	SlotId        int
	PaymentStatus string
	Patientdetail Patient
}
type GooglePatientDetails struct {
	GoogleID     string
	GoogleEmail  string
	AccessToken  string
	RefreshToken string
	TokenExpiry  string
}
