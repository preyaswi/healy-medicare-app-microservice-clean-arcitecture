package interfaces

import models "healy-apigateway/pkg/utils"

type ChatClient interface {
	GetChat(userid string, req models.ChatRequest) ([]models.TempMessage, error)
}
