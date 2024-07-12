package handler

import (
	"encoding/json"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"
	"log"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type PatientHandler struct {
	Grpc_client interfaces.PatientClient
	oauthConfig *oauth2.Config
}

func NewPatientHandler(PatientClient interfaces.PatientClient, cfg config.Config) *PatientHandler {
	return &PatientHandler{
		Grpc_client: PatientClient,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.GoogleClientId,
			ClientSecret: cfg.GoogleSecretId,
			RedirectURL:  cfg.RedirectURL,
			Scopes: []string{
				calendar.CalendarScope,
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

// GoogleLogin godoc
// @Summary Redirect to Google OAuth2 login //use browser
// @Description Redirects the user to Google's OAuth2 login page
// @Tags Patients
// @Accept json
// @Produce json
// @Success 302 {string} string "Redirecting to Google login"
// @Router /patient/login [get]
// @Security None
func (p *PatientHandler) GoogleLogin(c *fiber.Ctx) error {
	url := p.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// GoogleCallback godoc
// @Summary Handle Google OAuth2 callback //redirect from google login
// @Description Handles the callback from Google OAuth2 login
// @Tags Patients
// @Accept json
// @Produce json
// @Success 201 {object} models.GoogleUserInfo
// @Failure 400 {string} string "No code in query parameters"
// @Failure 500 {string} string "Failed to exchange token"
// @Router /google/redirect [get]
// @Security None
func (p *PatientHandler) GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		log.Printf("No code in query parameters")
		return c.Status(fiber.StatusBadRequest).SendString("No code in query parameters")
	}

	token, err := p.oauthConfig.Exchange(c.Context(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
	}

	// Retrieve user info using the token
	client := p.oauthConfig.Client(c.Context(), token)
	userInfoResponse, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Unable to retrieve user info: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to retrieve user info")
	}
	defer userInfoResponse.Body.Close()

	var userInfo struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Link          string `json:"link"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
		HD            string `json:"hd"`
	}

	if err := json.NewDecoder(userInfoResponse.Body).Decode(&userInfo); err != nil {
		log.Printf("Unable to parse user info: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to parse user info")
	}

	googleID := userInfo.ID
	googleEmail := userInfo.Email

	googleUser := models.GoogleUserInfo{
		ID:           googleID,
		Email:        googleEmail,
		Name:         userInfo.Name,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenExpiry:  token.Expiry.String(),
	}

	patient, err := p.Grpc_client.GoogleSignIn(googleUser.ID, googleUser.Email, googleUser.Name, googleUser.AccessToken, googleUser.RefreshToken, googleUser.TokenExpiry)
	if err != nil {
		errs := response.ClientResponse("error: Failed to authenticate with patient service", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)

	}
	success := response.ClientResponse("Patient created successfully", patient, nil)
	return c.Status(201).JSON(success)
}

// @Summary Get patient details
// @Description Retrieve details of a specific patient
// @Tags Patient Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /patient/profile [get]
func (p *PatientHandler) PatientDetails(c *fiber.Ctx) error {

	patientID := c.Locals("user_id").(string)
	patientDetails, err := p.Grpc_client.PatientDetails(patientID)
	if err != nil {
		errorRes := response.ClientResponse("failed to retrieve patient details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)

	}
	successRes := response.ClientResponse("patient Details", patientDetails, nil)
	return c.Status(200).JSON(successRes)

}
// @Summary Update patient details
// @Description Update the details of a patient
// @Tags Patient Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param patientDetails body models.PatientDetails true "Patient details to update"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /patient/profile [put]
func (p *PatientHandler) UpdatePatientDetails(c *fiber.Ctx) error {

	user_id := c.Locals("user_id").(string)

	var patient models.PatientDetails

	if err := c.BodyParser(&patient); err != nil {
		errs := response.ClientResponse("details are not in correct format", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	if err := validator.New().Struct(patient); err != nil {
		errs := response.ClientResponse("Constraints not satisfied", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}
	updatedDetails, err := p.Grpc_client.UpdatePatientDetails(patient, user_id)
	if err != nil {
		errorRes := response.ClientResponse("failed update user", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errorRes)
	}

	successRes := response.ClientResponse("Updated User Details", updatedDetails, nil)
	return c.Status(200).JSON(successRes)

}
// ListPatients handles listing all patients
// @Summary List Patients
// @Description List all patients
// @Tags Admin
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Security Bearer
// @Router /admin/dashboard/patients [get]
func (p *PatientHandler) ListPatients(c *fiber.Ctx) error {
	listedPatients, err := p.Grpc_client.ListPatients()
	if err != nil {
		errorRes := response.ClientResponse("failed retreiving list of patients", nil, err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorRes)
	}
	successRes := response.ClientResponse("retrived list of patients", listedPatients, nil)
	return c.Status(200).JSON(successRes)
}
