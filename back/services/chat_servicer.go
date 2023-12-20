package services

import "github.com/GrandOichii/messager-app/back/models"

type ChatServicer interface {
	Create(owner string, chatData *models.CreateChat) (*models.Chat, error)
}
