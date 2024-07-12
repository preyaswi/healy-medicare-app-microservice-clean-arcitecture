package interfaces

import "chatservice/pkg/models"

type ChatRepository interface {
	StoreFriendsChat(models.MessageReq) error
	UpdateReadAsMessage(string, string) error
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
}
