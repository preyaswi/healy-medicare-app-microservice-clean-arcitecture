package di

import (
	server "healy-admin/pkg/api"
	"healy-admin/pkg/api/service"
	"healy-admin/pkg/client"
	"healy-admin/pkg/config"
	"healy-admin/pkg/db"
	"healy-admin/pkg/repository"
	"healy-admin/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	doctorClient := client.NewdoctorClient(&cfg)
	patientClient := client.NewPatientClient(&cfg)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, doctorClient, patientClient)

	adminServiceServer := service.NewAdminServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
