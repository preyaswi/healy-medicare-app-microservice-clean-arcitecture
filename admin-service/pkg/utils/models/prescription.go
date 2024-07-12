package models

type PrescriptionRequest struct {
	DoctorID  int    `json:"doctor_id"`
	BookingID int    `json:"booking_id"`
	Medicine  string `json:"medicine" validate:"required"`
	Dosage    string `json:"dosage" validate:"required"`
	Notes     string `json:"notes"`
}
