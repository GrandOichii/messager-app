package services

import (
	"errors"

	"github.com/GrandOichii/messager-app/back/models"
)

type ChatService struct {
	ChatServicer

	chats []models.Chat
}

func (cs *ChatService) Create(owner string, chatData *models.CreateChat) (*models.Chat, error) {
	for _, chat := range cs.chats {
		if chat.HasParticipant(chatData.WithHandle) {
			// TODO return already existing chat?
			return nil, errors.New("chat with " + chatData.WithHandle + " already exists")
		}
	}

}
