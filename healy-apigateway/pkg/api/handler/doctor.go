package handler

import (
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type DoctorHandler struct {
	Grpc_Client interfaces.DoctorClient
}

func NewDoctorHandler(DoctorClient interfaces.DoctorClient) *DoctorHandler {
	return &DoctorHandler{
		Grpc_Client: DoctorClient,
	}
}
// @Summary Doctor Sign Up
// @Description Sign up a new doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param doctor body models.DoctorSignUp true "Doctor Sign Up Details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/signup [post]
// @Param doctor body models.DoctorSignUp true "Doctor Sign Up Details" Example({"full_name": "John Doe", "email": "john.doe@example.com", "phone_number": "1234567890", "password": "password123", "confirm_password": "password123", "specialization": "Cardiologist", "years_of_experience": 10, "license_number": "XYZ12345", "fees": 100})
func (d *DoctorHandler) DoctorSignUp(c *fiber.Ctx) error {
	var SignupDetail models.DoctorSignUp
	if err := c.BodyParser(&SignupDetail); err != nil {
		errs := response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(SignupDetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctor, err := d.Grpc_Client.DoctorSignUp(SignupDetail)
	if err != nil {
		errs := response.ClientResponse("Details not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctor created successfully", doctor, nil)
	return c.Status(201).JSON(success)
}
// @Summary Doctor Login
// @Description Log in an existing doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param login body models.DoctorLogin true "Doctor Login Details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/login [post]
// @Param login body models.DoctorLogin true "Doctor Login Details" Example({"email": "john.doe@example.com", "password": "password123"})
func (d *DoctorHandler) DoctorLogin(c *fiber.Ctx) error {
	var logindetail models.DoctorLogin
	if err := c.BodyParser(&logindetail); err != nil {
		errs := response.ClientResponse("logindetails are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	if err := validator.New().Struct(logindetail); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctor, err := d.Grpc_Client.DoctorLogin(logindetail)

	if err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctor logined succesfully", doctor, nil)
	return c.Status(201).JSON(success)
}
// @Summary List Doctors
// @Description List all doctors
// @Tags Admin
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Security Bearer
// @Router /admin/dashboard/doctors [get]
// @Summary Get all doctors' details
// @Description Retrieve the details of all doctors
// @Tags Patient
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /patient/doctor [get]
func (d *DoctorHandler) DoctorsDetails(c *fiber.Ctx) error {
	doctor, err := d.Grpc_Client.DoctorsDetails()
	if err != nil {
		errs := response.ClientResponse("couldnt fetch doctors data", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	success := response.ClientResponse("doctors data fetched succesfully", doctor, nil)
	return c.Status(201).JSON(success)
}
// @Summary Get individual doctor's details
// @Description Retrieve the details of a specific doctor by ID
// @Tags Patient
// @Accept json
// @Produce json
// @Security Bearer
// @Param doctor_id path string true "Doctor ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /patient/doctor/{doctor_id} [get]
func (d *DoctorHandler) IndividualDoctor(c *fiber.Ctx) error {
	doctorID := c.Params("doctor_id")
	doctor, err := d.Grpc_Client.IndividualDoctor(doctorID)
	if err != nil {
		errorRes := response.ClientResponse("couldn't fetch doctors data", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	success := response.ClientResponse("returned individual doctor data", doctor, nil)
	return c.Status(201).JSON(success)
}
// @Summary Get Doctor Profile
// @Description Retrieve doctor profile
// @Tags Doctor
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/profile [get]
func (d *DoctorHandler) DoctorProfile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	userID, err := strconv.Atoi(userId)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	doctorDetail, err := d.Grpc_Client.DoctorProfile(userID)
	if err != nil {
		errorRes := response.ClientResponse("failed to retrieve doctor details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)

	}
	successRes := response.ClientResponse("doctor Details", doctorDetail, nil)
	return c.Status(200).JSON(successRes)

}

// @Summary Rate a doctor
// @Description Rate a specific doctor
// @Tags Patient
// @Accept json
// @Produce json
// @Security Bearer
// @Param doctor_id path string true "Doctor ID"
// @Param rate body models.Rate true "Rate"
// @Success 202 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /patient/doctor/rate/{doctor_id} [post]
func (d *DoctorHandler) RateDoctor(c *fiber.Ctx) error {
	patientid := c.Locals("user_id").(string)
	doctor_id := c.Params("doctor_id")
	var rate models.Rate
	if err := c.BodyParser(&rate); err != nil {
		errs := response.ClientResponse("Details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(rate); err != nil {
		errs := response.ClientResponse("give rate between 1 to 5", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	rated, err := d.Grpc_Client.RateDoctor(patientid, doctor_id, rate)
	if err != nil {
		errs := response.ClientResponse("couldn't add the rating", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	successRes := response.ClientResponse("doctor Details", rated, nil)
	return c.Status(202).JSON(successRes)

}
// @Summary Update Doctor Profile
// @Description Update doctor profile
// @Tags Doctor
// @Accept json
// @Produce json
// @Security Bearer
// @Param profile body models.DoctorDetails true "Doctor Details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /doctor/profile [put]
// @Param profile body models.DoctorDetails true "Doctor Details" Example({"full_name": "John Doe", "email": "john.doe@example.com", "phone_number": "1234567890", "specialization": "Cardiologist", "years_of_experience": 10, "fees": 100})
func (d *DoctorHandler) UpdateDoctorProfile(c *fiber.Ctx) error {
	doctorid := c.Locals("user_id").(string)
	doctorId, err := strconv.Atoi(doctorid)
	if err != nil {
		errs := response.ClientResponse("couldn't convert string to int", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	var doctor models.DoctorDetails

	if err := c.BodyParser(&doctor); err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(doctor); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	res, err := d.Grpc_Client.UpdateDoctorProfile(doctorId, doctor)

	if err != nil {
		errorRes := response.ClientResponse("failed update doctor", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}
	successRes := response.ClientResponse("Updated doctor Details", res, nil)
	return c.Status(200).JSON(successRes)
}
