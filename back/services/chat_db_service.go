package services

import (
	"context"
	"errors"

	"github.com/GrandOichii/messager-app/back/constants"
	"github.com/GrandOichii/messager-app/back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CHATS_COLLECTION = "chats"
)

type ChatDBService struct {
	ChatServicer

	services *Services
	dbClient *mongo.Client
}

func NewChatDBService(client *mongo.Client, services *Services) *ChatDBService {
	return &ChatDBService{
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
	// TODO add the chat to the other user

	other, err := cs.services.UserServicer.ByHandle(chatData.WithHandle)
	if err != nil {
		return nil, err
	}

	cursor := cs.dbClient.Database(constants.DB_NAME).Collection(CHATS_COLLECTION).FindOne(context.TODO(), bson.D{
		{Key: "participants", Value: chatData.WithHandle},
	})

	err = cursor.Err()
	if err == nil {
		return nil, errors.New("chat with " + chatData.WithHandle + " already exists")
	} else if err != mongo.ErrNoDocuments {
		panic(err)
	}

	chat := &models.Chat{
		ParticipantHandles: []string{owner, other.Handle},
		Messages:           []*models.Message{},
	}

	inRes, err := cs.dbClient.Database(constants.DB_NAME).Collection(CHATS_COLLECTION).InsertOne(context.TODO(), chat)
	if err != nil {
		return nil, err
	}
	chatID := inRes.InsertedID.(primitive.ObjectID)

	for _, userHandle := range []string{
		owner,
		chatData.WithHandle,
	} {
		// * this will be here to horever mock me for my foolishness for searching for a bug for an hour and not simply having the ability to read my code
		// filter := bson.D{{Value: "handle", Key: userHandle}}
		filter := bson.D{{Key: "handle", Value: userHandle}}
		update := bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "chat_ids", Value: chatID},
			}},
		}
		uRes, err := cs.dbClient.Database(constants.DB_NAME).Collection(USERS_COLLECTION).UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
		if uRes.ModifiedCount == 0 {
			panic(errors.New("user with handle " + userHandle + " was not updated on creating chat with id " + chatID.String() + " because they were not found"))
		}
	}

	return chat, nil
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

	uRes, err := cs.dbClient.Database(constants.DB_NAME).Collection(CHATS_COLLECTION).UpdateByID(context.TODO(), chat.ID, bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "messages", Value: message},
		}},
	})

	if err != nil {
		panic(err)
	}

	if uRes.ModifiedCount == 0 {
		panic("failed to append message to chat with id " + chat.ID)
	}

	return message, nil
}
