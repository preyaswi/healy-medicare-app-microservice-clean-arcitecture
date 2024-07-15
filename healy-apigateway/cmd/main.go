package main

import (
	"healy-apigateway/cmd/docs"
	"healy-apigateway/pkg/config"
	"healy-apigateway/pkg/di"
	"healy-apigateway/pkg/logging"
	"log"
)

// @title Healy Medicare API
// @version 1.0.0
// @description This is the API documentation for the Healy Medicare application.
// @contact.name API Support
// @contact.email support@preyaswi.online
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	logging.Init()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load api gateway config", err)
	}
	// Update Swagger info
	docs.SwaggerInfo.Title = "Healy Medicare API"
	docs.SwaggerInfo.Description = "This is the API documentation for the Healy Medicare application."
	docs.SwaggerInfo.Version = "1.0.0"
    docs.SwaggerInfo.Host = "preyaswi.online"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	server, diErr := di.InitializeApi(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start(cfg)
	}
}
