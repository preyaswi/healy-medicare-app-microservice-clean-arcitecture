package usecase

import (
	"context"
	"errors"
	"fmt"
	clientinterface "healy-admin/pkg/client/interface"
	"healy-admin/pkg/config"
	"healy-admin/pkg/domain"
	"healy-admin/pkg/helper"
	interfaces "healy-admin/pkg/repository/interface"
	services "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/razorpay/razorpay-go"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type adminUseCase struct {
	adminRepository   interfaces.AdminRepository
	doctorRepository  clientinterface.NewDoctorClient
	patientRepository clientinterface.NewPatientClient
}

func NewAdminUseCase(repository interfaces.AdminRepository, doctorRepo clientinterface.NewDoctorClient, patientrepo clientinterface.NewPatientClient) services.AdminUseCase {
	return &adminUseCase{
		adminRepository:   repository,
		doctorRepository:  doctorRepo,
		patientRepository: patientrepo,
	}
}
func (ad *adminUseCase) AdminSignUp(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenAdmin{}, errors.New("admin with this email is already exists")
	}
	hashPassword, err := helper.PasswordHash(admin.Password)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error in hashing password")
	}
	admin.Password = hashPassword
	admindata, err := ad.adminRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not add the user")
	}
	tokenString, err := helper.GenerateTokenAdmin(admindata)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: admindata,
		Token: tokenString,
	}, nil
}

func (ad *adminUseCase) LoginHandler(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("email doesn't exist")
	}
	admindeatils, err := ad.adminRepository.FindAdminByEmail(admin)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}
	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(admindeatils.Password), []byte(admin.Password))
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("password not matching")
	}
	var adminDetailsResponse models.AdminDetailsResponse
	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &admindeatils)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil
}
func (ad *adminUseCase) AddToBooking(patientid string, doctorid int) error {
	ok, err := ad.doctorRepository.CheckDoctor(doctorid)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("doctor doesn't exist")
	}
	doctordetail, err := ad.doctorRepository.DoctorDetailforBooking(doctorid)
	if err != nil {
		return err
	}
	err = ad.adminRepository.AddToBooking(patientid, doctordetail)
	if err != nil {
		return err
	}
	return nil

}
func (ad *adminUseCase) CancelBooking(patientid string, bookingid int) error {
	booking, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return err
	}
	if booking.PatientId != patientid {
		return errors.New("unauthorized: patient ID does not match booking")
	}
	return ad.adminRepository.RemoveBooking(bookingid)

}
func (ad *adminUseCase) MakePaymentRazorpay(bookingid int) (domain.Booking, string, error) {
	cfg, _ := config.LoadConfig()
	paymentdetail, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	patientdetails,err:=ad.patientRepository.GetPatientByID(paymentdetail.PatientId)
	if err!=nil{
		return domain.Booking{}, "", err
	}
	paymentdetail.PatientId=patientdetails.Fullname
	fmt.Println(paymentdetail,"tyuiokn")
	client := razorpay.NewClient(cfg.KEY_ID_fOR_PAY, cfg.SECRET_KEY_FOR_PAY)
	data := map[string]interface{}{
		"amount":   int(paymentdetail.Fees) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Booking{}, "", err
	}
	RazorpayorderId := body["id"].(string)
	err = ad.adminRepository.AddRazorPayDetails(paymentdetail.BookingId, RazorpayorderId)
	if err != nil {
		return domain.Booking{}, "", err
	}
	return paymentdetail, RazorpayorderId, nil
}
func (ad *adminUseCase) VerifyPayment(booking_id int) error {
	status, err := ad.adminRepository.CheckPaymentStatus(booking_id)
	if err != nil {
		return err
	}
	status = strings.TrimSpace(strings.ToLower(status))
	if status == "not paid" {
		if err := ad.adminRepository.UpdatePaymentStatus(booking_id, "paid"); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("already paid")
	}
}

func (ad *adminUseCase) GetPaidPatients(doctor_id int) ([]models.BookedPatient, error) {
	bookings, err := ad.adminRepository.GetPaidBookingsByDoctorID(doctor_id)
	if err != nil {
		return nil, err
	}
	bookedPatients := make([]models.BookedPatient, len(bookings))
	var wg sync.WaitGroup
	mu := &sync.Mutex{} // Mutex to protect shared resources
	errors := make([]error, len(bookings))

	for i, booking := range bookings {
		wg.Add(1)
		go func(i int, booking domain.Booking) {
			defer wg.Done()
			patient, err := ad.patientRepository.GetPatientByID(booking.PatientId)
			if err != nil {
				mu.Lock()
				errors[i] = err
				mu.Unlock()
				return
			}
			mu.Lock()
			bookedPatients[i] = models.BookedPatient{
				BookingId:     int(booking.BookingId),
				SlotId: int(booking.Slot_id),
				PaymentStatus: booking.PaymentStatus,
				Patientdetail: patient,
			}
			mu.Unlock()
		}(i, booking)
	}

	wg.Wait()

	// Check for errors
	for _, err := range errors {
		if err != nil {
			return nil, fmt.Errorf("error fetching patient details: %v", err)
		}
	}

	return bookedPatients, nil
}
func (ad *adminUseCase) CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error) {
	bookingsdetails,err:=ad.adminRepository.GetBookingByID(prescription.BookingID)
	if err!=nil{
		return domain.Prescription{},err
	}
	if bookingsdetails.DoctorId!=uint(prescription.DoctorID){
		return domain.Prescription{},errors.New("the bookingid doctor and the request doctor id is not same")
	}
	if bookingsdetails.PaymentStatus!="paid"{
		return domain.Prescription{}, fmt.Errorf("patient has not paid the booking fee")
	}
	createdPrescription, err := ad.adminRepository.CreatePrescription(prescription)
	if err != nil {
		return domain.Prescription{}, fmt.Errorf("error creating prescription")
	}

	return createdPrescription, nil
}
func (ad *adminUseCase) SetDoctorAvailability(availabiity models.SetAvailability) (string, error) {
	status, err := ad.adminRepository.SetDoctorAvailability(availabiity)
	if err != nil {
		return "", err
	}
	return status, nil
}
func (ad *adminUseCase) GetDoctorAvailability(dotctorid int, date time.Time) ([]models.AvailableSlots, error) {
	availableSlots, err := ad.adminRepository.GetDoctorAvailability(dotctorid, date)
	if err != nil {
		return []models.AvailableSlots{}, err
	}
	return availableSlots, nil
}
func (ad *adminUseCase) BookSlot(patientid string, bookingid, slotid int) error {
	paid, err := ad.adminRepository.CheckPatientPayment(patientid, bookingid)
	if err != nil {
		return errors.New("error checking payment status")
	}

	if !paid {
		return errors.New("patient has not paid the booking fee")
	}
	slotAvailable, err := ad.adminRepository.CheckSlotAvailability(slotid)
	if err != nil {
		return err
	}
	if !slotAvailable {
		return errors.New("slot is already booked")
	}
	// Book the slot
	err = ad.adminRepository.BookSlot(bookingid, slotid)
	if err != nil {
		return err
	}

	// Mark slot as booked
	err = ad.adminRepository.MarkSlotAsBooked(slotid)
	if err != nil {
		return err
	}

	return nil

}
func (ad *adminUseCase) BookDoctor(patientid string, slotid int) (domain.Booking, string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return domain.Booking{}, "", errors.New("couldn't load razorpay credentials")
	}
	slotAvailable, err := ad.adminRepository.CheckSlotAvailability(slotid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	if !slotAvailable {
		return domain.Booking{}, "", errors.New("slot is already booked")
	}
	doctorid, err := ad.adminRepository.GetDoctorIdFromSlotId(slotid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	doctordetail, err := ad.doctorRepository.DoctorDetailforBooking(doctorid)
	fmt.Println(doctordetail, "this is the doctor details")
	if err != nil {
		return domain.Booking{}, "", err
	}
	bookingid, err := ad.adminRepository.AddDetailsToBooking(patientid, doctordetail, slotid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	
	client := razorpay.NewClient(cfg.KEY_ID_fOR_PAY, cfg.SECRET_KEY_FOR_PAY)
	data := map[string]interface{}{
		"amount":   int(doctordetail.Fees) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {

		return domain.Booking{}, "", err
	}
	RazorpayorderId := body["id"].(string)

	err = ad.adminRepository.AddRazorPayDetails(uint(bookingid), RazorpayorderId)
	if err != nil {
		return domain.Booking{}, "", err
	}
	paymentdetail, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	patientdetails,err:=ad.patientRepository.GetPatientByID(paymentdetail.PatientId)
	if err!=nil{
		return domain.Booking{}, "", err
	}
	paymentdetail.PatientId=patientdetails.Fullname
	return paymentdetail, RazorpayorderId, nil
}
func (ad *adminUseCase) VerifyandCalenderCreation(bookingid int, paymentid, razorid string) error {
	status, err := ad.adminRepository.CheckPaymentStatus(bookingid)
	if err != nil {
		return err
	}
	if status == "paid" {
		return errors.New("already paid")
	}
	if status == "initialized" {

		err = ad.adminRepository.UpdatePaymentDetails(bookingid, paymentid)
		if err != nil {
			return err
		}
		err = ad.adminRepository.UpdatePaymentStatus(bookingid, "paid")
		if err != nil {
			return err
		}
	}
	bookingdetails, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return err
	}
	err = ad.adminRepository.UpdateSlotAvailability(int(bookingdetails.Slot_id))
	if err != nil {
		return err
	}
	availability, err := ad.adminRepository.GetAvailabilityByID(int(bookingdetails.Slot_id))
	if err != nil {
		return err
	}
	patientGoogleDetails, err := ad.patientRepository.GetPatientGoogleDetailsByID(bookingdetails.PatientId)
	if err != nil {
		return err
	}
	doctordetail, err := ad.doctorRepository.DoctorDetailforBooking(int(bookingdetails.DoctorId))
	if err != nil {
		return err
	}
	eventID, err := ad.createGoogleCalendarEvent(patientGoogleDetails, availability, doctordetail.DoctorEmail)
	if err != nil {
		return err
	}
	err = ad.adminRepository.StoreEventDetails(domain.Event{
		BookingID:   uint(bookingid),
		PatientID:   bookingdetails.PatientId,
		EventID:     eventID,
		Summary:     "Doctor Appointment",
		Description: "Your scheduled appointment with the doctor",
		Start:       availability.StartTime,
		End:         availability.EndTime,
		GuestEmail:  bookingdetails.DoctorEmail,
	})
	if err != nil {
		return err
	}

	return nil

}
func (ad *adminUseCase) createGoogleCalendarEvent(patientDetails models.GooglePatientDetails, availability domain.Availability, doctoremail string) (string, error) {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", errors.New("couldn't load config for calender creation")
	}
	config := &oauth2.Config{
		ClientID:     cfg.GoogleClientId,
		ClientSecret: cfg.GoogleSecretId,
		Endpoint:     google.Endpoint,
		Scopes:       []string{calendar.CalendarScope},
	}

	// Parse token expiry
	tokenExpiry, err := time.Parse(time.RFC3339, patientDetails.TokenExpiry)
	if err != nil {
		// If parsing fails, try an alternative approach
		tokenExpiry, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", patientDetails.TokenExpiry)
		if err != nil {
			// If it still fails, use the current time as a fallback
			tokenExpiry = time.Now()
		}
	}

	token := &oauth2.Token{
		AccessToken:  patientDetails.AccessToken,
		RefreshToken: patientDetails.RefreshToken,
		TokenType:    "Bearer",
		Expiry:       tokenExpiry,
	}
	// Create a new token source with automatic token refresh
	tokenSource := config.TokenSource(ctx, token)

	// Create a client with the refreshing token source
	client := oauth2.NewClient(ctx, tokenSource)

	srv, err := calendar.New(client)
	if err != nil {
		return "", err
	}
	event := &calendar.Event{
		Summary:     "Doctor Appointment",
		Description: "Your scheduled appointment with the doctor",
		Start: &calendar.EventDateTime{
			DateTime: availability.StartTime.Format(time.RFC3339),
			TimeZone: "UTC", // e.g., "America/New_York"
		},
		End: &calendar.EventDateTime{
			DateTime: availability.EndTime.Format(time.RFC3339),
			TimeZone: "UTC", // e.g., "America/New_York"
		},
		Attendees: []*calendar.EventAttendee{
			{Email: doctoremail},
		},
	}

	createdEvent, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		return "", err
	}

	// If the token was refreshed, update it in the database
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", err
	}
	if newToken.AccessToken != token.AccessToken {
		err = ad.patientRepository.UpdatePatientGoogleToken(patientDetails.GoogleID, newToken.AccessToken, newToken.RefreshToken, newToken.Expiry.Format(time.RFC3339))
		if err != nil {
			return "", err
		}
	}

	return createdEvent.Id, nil

}
