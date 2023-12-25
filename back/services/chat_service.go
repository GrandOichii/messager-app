package services

import (
	"errors"

	"github.com/GrandOichii/messager-app/back/models"
)

type ChatService struct {
	ChatServicer

	services *Services
	chats    []*models.Chat
}

func NewChatService(services *Services) *ChatService {
	return &ChatService{
		chats:    []*models.Chat{},
		services: services,
	}
}

func (cs *ChatService) ByID(chatID string) (*models.Chat, error) {
	for _, chat := range cs.chats {
		if chat.ID.Hex() == chatID {
			return chat, nil
		}
	}
	return nil, errors.New("chat with ID " + chatID + " not found")
}

func (cs *ChatService) Create(owner string, chatData *models.CreateChat) (*models.Chat, error) {
	// TODO add the chat to the other user
	for _, chat := range cs.chats {
		if chat.HasParticipant(chatData.WithHandle) {
			// TODO return already existing chat?
			return nil, errors.New("chat with " + chatData.WithHandle + " already exists")
		}
	}

	other, err := cs.services.UserServicer.ByHandle(chatData.WithHandle)
	if err != nil {
		return nil, err
	}
	res := &models.Chat{
		ParticipantHandles: []string{owner, other.Handle},
		Messages:           []*models.Message{},
	}

	cs.chats = append(cs.chats, res)

	return res, nil
}

func (cs *ChatService) AddMessage(owner *models.User, chat *models.Chat, newMessage *models.PostMessage) (*models.Message, error) {
	message := &models.Message{
		Text:        newMessage.Text,
		OwnerHandle: owner.Handle,
	}

	err := message.CheckValid()
	if err != nil {
		return nil, err
	}

	chat.Messages = append(chat.Messages, message)

	return message, nil
}
