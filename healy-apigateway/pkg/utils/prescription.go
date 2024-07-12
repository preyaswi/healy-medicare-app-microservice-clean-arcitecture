package models

type Prescription struct {
	Medicine string `json:"medicine" validate:"required"`
	Dosage   string `json:"dosage" validate:"required"`
	Notes    string `json:"notes"`
}
// PrescriptionRequest model
// @Description Prescription request details
type PrescriptionRequest struct {
	DoctorID  int    `json:"doctor_id"`
	BookingID int    `json:"booking_id"`
	Medicine  string `json:"medicine" `
	Dosage    string `json:"dosage" `
	Notes     string `json:"notes"`
}
type CreatedPrescription struct {
	Id        int    `json:"id"`
	DoctorID  int    `json:"doctor_id"`
	PatientID string `json:"patient_id"`
	BookingID int    `json:"booking_id"`
	Medicine  string `json:"medicine" `
	Dosage    string `json:"dosage" `
	Notes     string `json:"notes"`
}
