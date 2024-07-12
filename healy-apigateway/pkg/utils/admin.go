package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}
type AdminSignUp struct {
	Firstname string `json:"firstname" binding:"required" validate:"required"`
	Lastname  string `json:"lastname" binding:"required" validate:"required"`
	Email     string `json:"email" binding:"required" validate:"required"`
	Password  string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"Email"`
}
type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin AdminDetailsResponse
	Token string
}

type Payment struct {
	PaymentId     uint
	PatientId     uint
	DoctorId      uint
	DoctorName    string
	Fees          uint64
	PaymentStatus string
}

// SetAvailability model
// @Description Set availability details
type SetAvailability struct {
	Date      string `json:"date"`      // e.g., "2024-06-20"
	StartTime string `json:"starttime"` // e.g., "09:00"
	EndTime   string `json:"endtime"`   // e.g., "17:00"
}
type GetAvailability struct {
	Slot_id   uint32
	Time      string // e.g., "09:00-09:30"
	Is_booked bool
}
