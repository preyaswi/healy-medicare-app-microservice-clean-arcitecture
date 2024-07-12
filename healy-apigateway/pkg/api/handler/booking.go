package handler

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/logging"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookingHandler struct {
	Grpc_Client interfaces.AdminClient
}

func NewBookingHandler(BookingClient interfaces.AdminClient) *BookingHandler {
	return &BookingHandler{
		Grpc_Client: BookingClient,
	}
}
// @Summary Get Booked Patients
// @Description Retrieve booked patients for the doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/patient [get]
func (b *BookingHandler) GetBookedPatients(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	patientdetails, err := b.Grpc_Client.GetPaidPatients(doctorId)
	if err != nil {
		errs := response.ClientResponse("couldn't fetch booked patients, please try again", nil, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("Successfully fetched booked patients", patientdetails, nil)
	return c.Status(fiber.StatusOK).JSON(success)
}
// @Summary Create Prescription
// @Description Create a prescription for a patient
// @Tags Doctor
// @Accept json
// @Produce json
// @Security Bearer
// @Param prescription body models.PrescriptionRequest true "Prescription Details"
// @Param booking_id query string true "Booking ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/patient/prescription [post]
func (b *BookingHandler) CreatePrescription(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	bookingIDStr := c.Query("booking_id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert booking ID string to int", nil, err.Error()))
	}

	var prescription models.PrescriptionRequest
	if err := c.BodyParser(&prescription); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Details are not in correct format", nil, err.Error()))
	}

	prescription.DoctorID = doctorId
	prescription.BookingID = bookingID

	if err := validator.New().Struct(prescription); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Constraints not satisfied", nil, err.Error()))
	}

	createdPrescription, err := b.Grpc_Client.CreatePrescription(prescription)
	if err != nil {
		errorRes := response.ClientResponse("failed to create prescription", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("prescription created", createdPrescription, nil)
	return c.Status(200).JSON(successRes)

}
// @Summary Set Doctor Availability
// @Description Set availability for the doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Security Bearer
// @Param availability body models.SetAvailability true "Availability Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/profile/availability [post]
// @Param availability body models.SetAvailability true "Availability Details" Example({"date": "2024-07-20", "starttime": "09:00", "endtime": "17:00"})
func (b *BookingHandler) SetDoctorAvailability(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errorRes := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	var setreq models.SetAvailability
	if err := c.BodyParser(&setreq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Details are not in correct format", nil, err.Error()))
	}
	res, err := b.Grpc_Client.SetDoctorAvailability(setreq, doctorId)
	if err != nil {
		errorRes := response.ClientResponse("failed to set availability", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("slot availability created created", res, nil)
	return c.Status(200).JSON(successRes)

}
// @Summary Get doctor's slot availability
// @Description Retrieve the available slots for a doctor on a given date
// @Tags Patient
// @Accept json
// @Produce json
// @Security Bearer
// @Param doctor_id query int true "Doctor ID"
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /patient/doctor/availability [get]
func (b *BookingHandler) GetDoctorSlotAvailability(c *fiber.Ctx) error {
	doctorId, err := strconv.Atoi(c.Query("doctor_id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert doctor ID string to int", nil, err.Error()))
	}
	date := c.Query("date")
	res, err := b.Grpc_Client.GetDoctorAvailability(doctorId, date)
	if err != nil {
		errorRes := response.ClientResponse("failed to get doctor's availability", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("listed doctor's availabile slots", res, nil)
	return c.Status(200).JSON(successRes)
}
// @Summary Book a doctor
// @Description Book a doctor for a specific slot,//user browser
// @Tags Patient
// @Accept json
// @Produce html
// @Param slot_id query int true "Slot ID"
// @Param patient_id query string true "Patient ID"
// @Success 200 {string} string "HTML page for payment"
// @Failure 500 {object} response.Response{}
// @Router /patient/bookdoctor [get]
func (b *BookingHandler) BookDoctor(c *fiber.Ctx) error {
	slotIdStr := c.Query("slot_id")
	slotID, err := strconv.Atoi(slotIdStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Cannot convert slot ID string to int", nil, err.Error()))
	}

	patientId := c.Query("patient_id")
	bookingdetails, razorId, err := b.Grpc_Client.BookDoctor(patientId, slotID)
	if err != nil {
		errorRes := response.ClientResponse("Couldn't book Doctor", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	fmt.Println(bookingdetails, "paymentdetails")
	fmt.Println(razorId, "razorid")
	fmt.Println(bookingdetails.PatientName,"dfghjk")
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{
		"final_price": bookingdetails.Fees * 100,
		"razor_id":    razorId,
		"user_id":     bookingdetails.PatientName,
		"order_id":    bookingdetails.BookingId,
		"user_email":  bookingdetails.DoctorEmail,
		"total":       int(bookingdetails.Fees),
	})
}

// @Summary Verify payment and create calendar entry  //redirect from book doctor
// @Description Verify the payment and create a calendar entry for the booking
// @Tags Patient
// @Accept json
// @Produce json
// @Param booking_id query int true "Booking ID"
// @Param payment_id query string true "Payment ID"
// @Param razor_id query string true "Razor ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /payment_success [get]
func (b *BookingHandler) VerifyandCalenderCreation(c *fiber.Ctx) error {
	logger := logging.Logger().WithField("function", "VerifyandCalenderCreation")
	logger.Info("Payment success endpoint hit")

	bookingIdStr := c.Query("booking_id")
	paymentId := c.Query("payment_id")
	razorId := c.Query("razor_id")

	logger = logger.WithFields(logrus.Fields{
		"booking_id": bookingIdStr,
		"payment_id": paymentId,
		"razor_id":   razorId,
	})
	logger.Info("Received payment verification request")

	bookingId, err := strconv.Atoi(bookingIdStr)
	if err != nil {
		logger.WithError(err).Error("Failed to convert booking ID to int")
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Invalid booking ID", nil, "Booking ID must be a valid integer"))
	}

	logger.Info("Calling gRPC VerifyandCalenderCreation")
	err = b.Grpc_Client.VerifyandCalenderCreation(bookingId, paymentId, razorId)
	if err != nil {
		logger.WithError(err).Error("gRPC VerifyandCalenderCreation failed")
		fmt.Println("gRPC error:", err.Error()) // Print the full error message

		// Determine if this is a client error or a server error
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "invalid") {
			return c.Status(http.StatusBadRequest).JSON(response.ClientResponse("Verification failed", nil, err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse("Internal server error", nil, "An unexpected error occurred"))
	}

	logger.Info("Slot booked successfully")
	return c.Status(http.StatusOK).JSON(response.ClientResponse("Slot booked on the calendar", nil, nil))
}
