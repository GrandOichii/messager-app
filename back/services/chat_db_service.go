package services

import (
	"context"
	"errors"

	"github.com/GrandOichii/messager-app/back/constants"
	"github.com/GrandOichii/messager-app/back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CHATS_COLLECTION = "chats"
)

type ChatDBService struct {
	ChatServicer

	services *Services
	dbClient *mongo.Client
	chats    []*models.Chat
}

func NewChatDBService(client *mongo.Client, services *Services) *ChatDBService {
	return &ChatDBService{
		chats:    []*models.Chat{},
		services: services,
		dbClient: client,
	}
}

func (cs *ChatDBService) ByID(chatID string) (*models.Chat, error) {
	// TODO add message pages
	cursor := cs.dbClient.Database(constants.DB_NAME).Collection(CHATS_COLLECTION).FindOne(context.TODO(), bson.D{
		{Key: "_id", Value: chatID},
	})

	var result models.Chat
	err := cursor.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no chat with id " + chatID)
		}
		panic(err)
	}
	return &result, nil
}

func (cs *ChatDBService) Create(owner string, chatData *models.CreateChat) (*models.Chat, error) {

	other, err := cs.services.UserServicer.ByHandle(chatData.WithHandle)
	if err != nil {
		return nil, err
	}

	// for _, chat := range cs.chats {
	// 	if chat.HasParticipant(chatData.WithHandle) {
	// 		// TODO return already existing chat?
	// 		return nil, errors.New("chat with " + chatData.WithHandle + " already exists")
	// 	}
	// }

	res := &models.Chat{
		ParticipantHandles: []string{owner, other.Handle},
		Messages:           []*models.Message{},
	}

	_, err = cs.dbClient.Database(constants.DB_NAME).Collection(CHATS_COLLECTION).InsertOne(context.TODO(), res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cs *ChatDBService) AddMessage(owner *models.User, chat *models.Chat, newMessage *models.PostMessage) (*models.Message, error) {
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
