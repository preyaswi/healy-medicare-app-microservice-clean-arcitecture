package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)
// @Summary Patient Routes
// @Description Group of routes for Patient operations
func PatientRoutes(route *fiber.App, patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler, bookingHandler *handler.BookingHandler,) {
	route.Get("/google/redirect", middleware.LoggingMiddleware(patientHandler.GoogleCallback))
	route.Get("/payment_success", middleware.LoggingMiddleware(bookingHandler.VerifyandCalenderCreation))

	patient := route.Group("/patient")
	patient.Get("login", middleware.LoggingMiddleware(patientHandler.GoogleLogin))
	patient.Get("/bookdoctor", middleware.LoggingMiddleware(bookingHandler.BookDoctor))
	patient.Use(middleware.UserAuthMiddleware())

	profile := patient.Group("/profile")
	profile.Get("", middleware.LoggingMiddleware(patientHandler.PatientDetails))
	profile.Put("", middleware.LoggingMiddleware(patientHandler.UpdatePatientDetails))

	doctor := patient.Group("/doctor")
	doctor.Get("availability", middleware.LoggingMiddleware(bookingHandler.GetDoctorSlotAvailability))
	doctor.Get("", middleware.LoggingMiddleware(doctorHandler.DoctorsDetails))
	doctor.Get("/:doctor_id", middleware.LoggingMiddleware(doctorHandler.IndividualDoctor))
	doctor.Post("/rate/:doctor_id", middleware.LoggingMiddleware(doctorHandler.RateDoctor))

}
