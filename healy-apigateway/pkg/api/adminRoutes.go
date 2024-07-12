package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(route *fiber.App, adminHandler *handler.AdminHandler, patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler) {
	admin := route.Group("/admin")
	admin.Post("/signup", middleware.LoggingMiddleware(adminHandler.AdminSignUp))
	admin.Post("/login", middleware.LoggingMiddleware(adminHandler.LoginHandler))
	admin.Use(middleware.AdminAuthMiddleware())
	{
		dashboard := admin.Group("/dashboard")
		patient := dashboard.Group("/patients")
		{
			patient.Get("", middleware.LoggingMiddleware(patientHandler.ListPatients))
		}
		doctor := dashboard.Group("/doctors")
		{
			doctor.Get("", middleware.LoggingMiddleware(doctorHandler.DoctorsDetails))
		}
	}
}
