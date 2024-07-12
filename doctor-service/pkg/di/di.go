package di

import (
	server "doctor-service/pkg/api"
	"doctor-service/pkg/api/service"
	"doctor-service/pkg/config"
	"doctor-service/pkg/db"
	"doctor-service/pkg/repository"
	"doctor-service/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectToDb(cfg)
	if err != nil {
		return nil, err
	}
	doctorRepository := repository.NewDoctorRepository(gormDB)
	doctorUseCase := usecase.NewDoctorUseCase(doctorRepository)

	doctorService := service.NewDoctorServer(doctorUseCase)
	doctorserver, err := server.NewGRPCServer(cfg, doctorService)
	if err != nil {
		return &server.Server{}, err
	}
	return doctorserver, nil
}
