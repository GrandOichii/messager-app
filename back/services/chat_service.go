package services

import (
	"errors"

	"github.com/GrandOichii/messager-app/back/models"
)

type ChatService struct {
	ChatServicer

	chats []*models.Chat
}

func NewChatService() *ChatService {
	return &ChatService{
		chats: []*models.Chat{},
	}
}

func (cs *ChatService) ByID(chatID string) (*models.Chat, error) {
	for _, chat := range cs.chats {
		if chat.ID == chatID {
			return chat, nil
		}
	}
	return nil, errors.New("chat with ID " + chatID + " not found")
}

func (cs *ChatService) Create(owner string, chatData *models.CreateChat) (*models.Chat, error) {
	for _, chat := range cs.chats {
		if chat.HasParticipant(chatData.WithHandle) {
			// TODO return already existing chat?
			return nil, errors.New("chat with " + chatData.WithHandle + " already exists")
		}
	}

	res := &models.Chat{
		ParticipantHandles: []string{owner, chatData.WithHandle},
		Messages:           []*models.Message{},
	}

	cs.chats = append(cs.chats, res)

	return res, nil
}

func (cs *ChatService) AddMessage(owner *models.User, chat *models.Chat, newMessage *models.PostMessage) (*models.Message, error) {
	// TODO add more complex messages
	message := &models.Message{
		Text:        newMessage.Text,
		OwnerHandle: owner.Handle,
	}

	chat.Messages = append(chat.Messages, message)

	return message, nil
}
