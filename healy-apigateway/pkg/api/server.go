package server

import (
	"healy-apigateway/pkg/api/handler"
	"healy-apigateway/pkg/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(patientHandler *handler.PatientHandler, doctorHandler *handler.DoctorHandler, adminHandler *handler.AdminHandler, bookingHandler *handler.BookingHandler, chatHandler *handler.ChatHandler) *ServerHTTP {
	engine := html.New("./template", ".html")
	route := fiber.New(fiber.Config{
		Views: engine,
	})
	route.Get("/swagger/*", swagger.HandlerDefault)
	route.Use(logger.New())

	route.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))


	DoctorRoutes(route, doctorHandler, patientHandler, bookingHandler)
	PatientRoutes(route, patientHandler, doctorHandler, adminHandler, bookingHandler)
	AdminRoutes(route, adminHandler, patientHandler, doctorHandler)
	ChatRoute(route, chatHandler)
	return &ServerHTTP{
		engine: route,
	}
}
func (s *ServerHTTP) Start(cfg config.Config) {

	log.Printf("starting server on :8000")
	err := s.engine.Listen(cfg.Port)
	if err != nil {
		log.Fatalf("error while starting the server: %v", err)
	}
}
