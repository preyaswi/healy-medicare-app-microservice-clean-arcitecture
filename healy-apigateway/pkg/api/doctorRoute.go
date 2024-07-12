package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)
// @Summary Doctor Routes
// @Description Group of routes for doctor operations
func DoctorRoutes(route *fiber.App, doctorHandler *handler.DoctorHandler, patientHandler *handler.PatientHandler, bookingHandler *handler.BookingHandler) {
	doctor := route.Group("/doctor")
	doctor.Post("/signup", middleware.LoggingMiddleware(doctorHandler.DoctorSignUp))
	doctor.Post("/login", middleware.LoggingMiddleware(doctorHandler.DoctorLogin))
	
	doctor.Use(middleware.UserAuthMiddleware())
	{
		profile := doctor.Group("/profile")
		profile.Get("", middleware.LoggingMiddleware(doctorHandler.DoctorProfile))
		profile.Put("", middleware.LoggingMiddleware(doctorHandler.UpdateDoctorProfile))
		profile.Post("/availability", bookingHandler.SetDoctorAvailability)
	}
	patient := doctor.Group("/patient")
	{
		patient.Get("", middleware.LoggingMiddleware(bookingHandler.GetBookedPatients))
		patient.Post("/prescription", middleware.LoggingMiddleware(bookingHandler.CreatePrescription))
	}

}
