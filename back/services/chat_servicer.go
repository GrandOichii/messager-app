package services

import "github.com/GrandOichii/messager-app/back/models"

type ChatServicer interface {
	ByID(chatID string) (*models.Chat, error)

	Create(owner string, chatData *models.CreateChat) (*models.Chat, error)
	AddMessage(owner *models.User, chat *models.Chat, newMessage *models.PostMessage) (*models.Message, error)
}
