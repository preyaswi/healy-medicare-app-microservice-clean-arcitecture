package di

import (
	server "chatservice/pkg/api"
	"chatservice/pkg/api/service"
	"chatservice/pkg/config"
	"chatservice/pkg/db"
	"chatservice/pkg/repository"
	"chatservice/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	db, err := db.ConnectToDatabase(&cfg)
	if err != nil {
		return nil, err
	}
	chatRepo := repository.NewChatRepository(db)
	chatUsecase := usecase.NewChatUseCase(chatRepo)
	chatServer := service.NewChatServer(chatUsecase)
	grpcServer, err := server.NewGRPCServer(cfg, chatServer)
	if err != nil {
		return &server.Server{}, err
	}
	go chatUsecase.MessageConsumer()
	return grpcServer, nil

}
