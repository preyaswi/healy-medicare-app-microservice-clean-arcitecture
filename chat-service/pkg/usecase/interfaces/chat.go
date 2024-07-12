package interfaces

import "chatservice/pkg/models"

type ChatUseCase interface {
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
	MessageConsumer()
}
