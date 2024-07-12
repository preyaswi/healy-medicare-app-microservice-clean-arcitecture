package main

import (
	"doctor-service/pkg/config"
	"doctor-service/pkg/di"
	"log"
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
