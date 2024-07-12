package di

import (
	server "patient-service/pkg/api"
	"patient-service/pkg/api/service"
	"patient-service/pkg/config"
	"patient-service/pkg/db"
	"patient-service/pkg/repository"
	"patient-service/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectToDb(cfg)
	if err != nil {
		return nil, err
	}
	patientRepository := repository.NewPatientRepository(gormDB)
	patientUseCase := usecase.NewPatientUseCase(patientRepository)

	patientServiceServer := service.NewPatientServer(patientUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, patientServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
