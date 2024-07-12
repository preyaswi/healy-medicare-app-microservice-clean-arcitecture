package main

import (
	"log"
	"patient-service/pkg/config"
	"patient-service/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	server, err := di.InitializeApi(config)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	} else {
		server.Start()
	}
}
