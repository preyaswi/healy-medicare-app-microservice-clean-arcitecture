package repository

import (
	"errors"
	"fmt"
	"healy-admin/pkg/domain"
	interfaces "healy-admin/pkg/repository/interface"
	"healy-admin/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse
	if err := ad.DB.Raw("INSERT INTO admins (firstname,lastname,email,password) VALUES (?, ?, ? ,?) RETURNING id,firstname,lastname,email", adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	return model, nil
}
func (ad *adminRepository) CheckAdminExistsByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where(&domain.Admin{Email: email}).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *adminRepository) FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error) {
	var user models.AdminSignUp
	err := ad.DB.Raw("SELECT * FROM admins WHERE email=? ", admin.Email).Scan(&user).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}

func (ad *adminRepository) AddToBooking(patientid string, doctordetail models.BookingDoctorDetails) error {
	err := ad.DB.Exec("insert into bookings(patient_id,doctor_id,doctor_name,doctor_email,fees)values(?,?,?,?,?)", patientid, doctordetail.Doctorid, doctordetail.DoctorName, doctordetail.DoctorEmail, doctordetail.Fees).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) AddDetailsToBooking(patientid string, doctordetail models.BookingDoctorDetails, slotid int) (int, error) {
	booking := domain.Booking{
		PatientId:   patientid,
		DoctorId:    doctordetail.Doctorid,
		DoctorName:  doctordetail.DoctorName,
		DoctorEmail: doctordetail.DoctorEmail,
		Fees:        doctordetail.Fees,
		Slot_id:     uint(slotid),
	}

	// Create the booking record
	result := ad.DB.Create(&booking)

	// Check for errors
	if result.Error != nil {
		return 0, result.Error
	}

	// Return the ID of the newly created booking
	return int(booking.BookingId), nil
}
func (ad *adminRepository) GetBookingByID(bookingid int) (domain.Booking, error) {
	var booking domain.Booking
	err := ad.DB.Where("booking_id=?", bookingid).First(&booking).Error
	if err != nil {
		return domain.Booking{}, err
	}
	return booking, nil
}
func (ad *adminRepository) UpdateSlotAvailability(slotid int) error {
	slot := &domain.Availability{}

	// Find the slot by ID and update the IsBooked field to true
	err := ad.DB.Model(slot).Where("id = ?", slotid).Update("is_booked", true).Error
	if err != nil {
		return err
	}

	return nil
}
func (ar *adminRepository) GetAvailabilityByID(slotID int) (domain.Availability, error) {
	var availability domain.Availability
	err := ar.DB.First(&availability, slotID).Error
	return availability, err
}
func (ar *adminRepository) StoreEventDetails(event domain.Event) error {
	return ar.DB.Create(&event).Error
}
func (ad *adminRepository) RemoveBooking(bookingID int) error {
	err := ad.DB.Where("booking_id=?", bookingID).Delete(&domain.Booking{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) AddRazorPayDetails(booking_id uint, razorPaypaymentID string) error {
	err := ad.DB.Exec("insert into razer_pays (booking_id,razor_id) values (?,?)", booking_id, razorPaypaymentID).Error
	if err != nil {
		return err
	}
	err = ad.DB.Model(&domain.Booking{}).Where("booking_id = ?", booking_id).Update("payment_status", "initialized").Error
	if err != nil {
		return err
	}
	return nil

}
func (ad *adminRepository) CheckPaymentStatus(bookingid int) (string, error) {
	var payment domain.Booking
	if err := ad.DB.First(&payment, bookingid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("payment not found")
		}
		return "", err
	}
	return payment.PaymentStatus, nil
}
func (ad *adminRepository) UpdatePaymentStatus(booking_id int, status string) error {
	err := ad.DB.Model(&domain.Booking{}).Where("booking_id = ?", booking_id).Update("payment_status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) GetPaidBookingsByDoctorID(doctorId int) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := ad.DB.Where("doctor_id = ? AND payment_status = ?", doctorId, "paid").Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (ad *adminRepository) CheckPatientPayment(patientid string, bookingid int) (bool, error) {
	var booking domain.Booking
	err := ad.DB.Where("patient_id = ? AND booking_id=? AND payment_status = ?", patientid, bookingid, "paid").First(&booking).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("error checking payment status")
	}

	return true, nil
}

func (ad *adminRepository) CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error) {
	prescriptionModel := domain.Prescription{
		BookingID: uint(prescription.BookingID),
		DoctorID:  uint(prescription.DoctorID),
		Medicine:  prescription.Medicine,
		Dosage:    prescription.Dosage,
		Notes:     prescription.Notes,
	}

	if err := ad.DB.Create(&prescriptionModel).Error; err != nil {
		return domain.Prescription{}, fmt.Errorf("error saving prescription")
	}

	// Retrieve the created prescription with all fields populated
	var createdPrescription domain.Prescription
	if err := ad.DB.First(&createdPrescription, prescriptionModel.ID).Error; err != nil {
		return domain.Prescription{}, fmt.Errorf("error retrieving created prescription")
	}

	return createdPrescription, nil
}
func (ad *adminRepository) SetDoctorAvailability(availability models.SetAvailability) (string, error) {
	fmt.Println(availability, "availability")
	var slots []domain.Availability
	currentTime := availability.StartTime
	for currentTime.Before(availability.EndTime) {
		// Define the end time for the current slot
		slotEndTime := currentTime.Add(30 * time.Minute)

		// Combine date and time
		startDateTime := time.Date(
			availability.Date.Year(),
			availability.Date.Month(),
			availability.Date.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
			currentTime.Second(),
			currentTime.Nanosecond(),
			availability.Date.Location(),
		)

		endDateTime := time.Date(
			availability.Date.Year(),
			availability.Date.Month(),
			availability.Date.Day(),
			slotEndTime.Hour(),
			slotEndTime.Minute(),
			slotEndTime.Second(),
			slotEndTime.Nanosecond(),
			availability.Date.Location(),
		)

		// Check for overlapping slots
		var existingSlots []domain.Availability
		err := ad.DB.Where("doctor_id = ? AND date = ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?) OR (start_time >= ? AND end_time <= ?))",
			availability.DoctorId, availability.Date, endDateTime, startDateTime, startDateTime, endDateTime, startDateTime, endDateTime).Find(&existingSlots).Error

		if err != nil {
			return "", err
		}

		// If no overlapping slots found, add the new slot to the list
		if len(existingSlots) == 0 {
			slots = append(slots, domain.Availability{
				DoctorID:  uint(availability.DoctorId),
				Date:      availability.Date,
				StartTime: startDateTime,
				EndTime:   endDateTime,
				IsBooked:  false,
			})
		}

		// Move to the next 30-minute slot
		currentTime = slotEndTime
	}

	// Check if there are slots to be inserted
	if len(slots) == 0 {
		fmt.Println(availability, "the availability")
		return "", errors.New("no available slots to be added")
	}

	// Save the non-overlapping slots to the database
	if err := ad.DB.Create(&slots).Error; err != nil {
		return "", err
	}
	return "success", nil
}
func (ad *adminRepository) GetDoctorAvailability(doctor_id int, date time.Time) ([]models.AvailableSlots, error) {
	var slots []domain.Availability
	if err := ad.DB.Where("doctor_id = ? AND date = ?", doctor_id, date).Find(&slots).Error; err != nil {
		return []models.AvailableSlots{}, err
	}
	fmt.Println(slots, "the slots when getting")
	var newslots []models.AvailableSlots
	for _, slot := range slots {
		newslots = append(newslots, models.AvailableSlots{
			Slot_id:  int(slot.ID),
			Time:     fmt.Sprintf("%s-%s", slot.StartTime.Format("15:04"), slot.EndTime.Format("15:04")),
			IsBooked: slot.IsBooked,
		})
	}
	return newslots, nil
}
func (ad *adminRepository) CheckSlotAvailability(slotid int) (bool, error) {
	var slot domain.Availability
	err := ad.DB.Where("id = ? AND is_booked = ?", slotid, false).First(&slot).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (ad *adminRepository) BookSlot(bookingid, slotid int) error {
	err := ad.DB.Exec("UPDATE bookings SET slot_id = ? WHERE booking_id = ?", slotid, bookingid).Error
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) MarkSlotAsBooked(slotid int) error {
	err := ad.DB.Model(&domain.Availability{}).Where("id = ?", slotid).Update("is_booked", true).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) GetDoctorIdFromSlotId(slotid int) (int, error) {
	var availability domain.Availability

	result := ad.DB.Model(&domain.Availability{}).Where("id = ?", slotid).First(&availability)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(availability.DoctorID), nil

}
func (ad *adminRepository) UpdatePaymentDetails(bookingid int, paymentid string) error {
	err := ad.DB.Exec("update razer_pays set payment_id = ? where booking_id = ?", paymentid, bookingid).Error
	if err != nil {
		return err
	}
	return nil
}
