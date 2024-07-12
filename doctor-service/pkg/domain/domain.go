package domain

type Doctor struct {
	ID                uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName          string `json:"full_name" gorm:"not null"`
	Email             string `json:"email" gorm:"unique"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password"`
	Specialization    string `json:"specialization"`
	YearsOfExperience int32  `json:"years_of_experience"`
	LicenseNumber     string `json:"license_number" gorm:"unique"`
	Fees              int64  `json:"fees" gorm:"not null"`
}
type Review struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	DoctorID  uint   `json:"doctor_id" gorm:"not null;uniqueIndex:unique_doctor_patient"`
	PatientID string `json:"patient_id" gorm:"not null;uniqueIndex:unique_doctor_patient"`
	Rating    int    `json:"rating" gorm:"not null"`
}
