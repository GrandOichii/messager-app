package services

import (
	"context"
	"errors"

	"github.com/GrandOichii/messager-app/back/constants"
	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/security"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USERS_COLLECTION = "users"
)

type UserDBService struct {
	UserServicer

	dbClient *mongo.Client
}

func NewUserDBService(dbClient *mongo.Client) *UserDBService {
	return &UserDBService{
		dbClient: dbClient,
	}
}

func toGetUserArr(arr []*models.User) []*models.GetUser {
	res := make([]*models.GetUser, len(arr))
	for i, user := range arr {
		res[i] = user.ToGetUser()
	}
	return res
}

func (us *UserDBService) All() ([]*models.GetUser, error) {
	cursor, err := us.dbClient.Database(constants.DB_NAME).Collection(USERS_COLLECTION).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return toGetUserArr(users), nil
}

func (us *UserDBService) ByHandle(handle string) (*models.User, error) {
	cursor := us.dbClient.Database(constants.DB_NAME).Collection(USERS_COLLECTION).FindOne(context.TODO(), bson.D{
		{Key: "handle", Value: handle},
	})

	var result models.User
	err := cursor.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no user with handle " + handle)
		}
		panic(err)
	}
	return &result, nil
}

func (us *UserDBService) Register(newUser *models.CreateUser) (*models.GetUser, error) {
	_, err := us.ByHandle(newUser.Handle)

	if err == nil {
		return nil, errors.New("user with handle " + newUser.Handle + " already exists")
	}

	passHash, err := security.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	res := &models.User{
		Handle: newUser.Handle,
		// TODO hash email
		EmailHash:    newUser.Email,
		PasswordHash: passHash,
		ChatIDs:      []string{},
	}

	_, err = us.dbClient.Database(constants.DB_NAME).Collection(USERS_COLLECTION).InsertOne(context.TODO(), res)
	if err != nil {
		return nil, err
	}

	return res.ToGetUser(), nil
}

func (us *UserDBService) Login(userData *models.LoginUser) (*models.User, error) {
	loginFailedErr := errors.New("failed to login")

	// TODO use hashed email
	cursor := us.dbClient.Database(constants.DB_NAME).Collection(USERS_COLLECTION).FindOne(context.TODO(), bson.D{
		{Key: "emailhash", Value: userData.Email},
	})

	var result models.User
	err := cursor.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, loginFailedErr
		}
		panic(err)
	}

	if !security.CheckPasswordHash(userData.Password, result.PasswordHash) {
		return nil, loginFailedErr
	}

	return &result, nil
}

func (us *UserDBService) GetChatIDs(handle string) ([]string, error) {
	user, err := us.ByHandle(handle)
	if err != nil {
		return nil, err
	}

	return user.ChatIDs, nil
}
