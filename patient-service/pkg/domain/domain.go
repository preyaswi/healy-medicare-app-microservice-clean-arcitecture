package domain

import (
	"time"
)

type Patient struct {
	ID            string `gorm:"primary_key"`
	GoogleId      string `json:"googleid" gorm:"unique;not null"`
	Fullname      string `json:"fullname" gorm:"validate:required"`
	Email         string `json:"email" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Contactnumber string `json:"contactnumber" gorm:"validate:required"`
	AccessToken   string `json:"accesstoken"`
	RefreshToken  string `json:"refreshtoken"`
	TokenExpiry   string `json:"tokenexpiry"`
}
type Prescription struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	PatientID  uint      `json:"patient_id" gorm:"not null"`
	DoctorID   uint      `json:"doctor_id" gorm:"not null"`
	DoctorName string    `json:"doctor_name" gorm:"not null"`
	Medicine   string    `json:"medicine" gorm:"not null"`
	Dosage     string    `json:"dosage" gorm:"not null"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
}
