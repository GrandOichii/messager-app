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
	COLLECTION = "users"
)

type UserDBService struct {
	UserServicer

	dbClient *mongo.Client
	users    []*models.User
}

func NewUserDBService(dbClient *mongo.Client) *UserDBService {
	return &UserDBService{
		dbClient: dbClient,
		users:    []*models.User{},
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
	cursor, err := us.dbClient.Database(constants.DB_NAME).Collection(COLLECTION).Find(context.TODO(), bson.D{})
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
	for _, user := range us.users {
		if user.Handle == handle {
			return user, nil
		}
	}
	return nil, errors.New("no user with handle " + handle)
}

func (us *UserDBService) Register(newUser *models.CreateUser) (*models.GetUser, error) {
	cursor, err := us.dbClient.Database(constants.DB_NAME).Collection(COLLECTION).Find(context.TODO(), bson.D{
		{Key: "handle", Value: newUser.Handle},
	})
	if err != nil {
		return nil, err
	}

	var users []*models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("User with handle " + newUser.Handle + " already exists")
	}

	// if err != nil {
	// 	return nil, err
	// }

	res := &models.User{
		Handle: newUser.Handle,
		// TODO hash email
		EmailHash: newUser.Email,
		// TODO hash password
		PasswordHash: newUser.Password,
	}

	_, err = us.dbClient.Database(constants.DB_NAME).Collection(COLLECTION).InsertOne(context.TODO(), res)
	if err != nil {
		return nil, err
	}

	return res.ToGetUser(), nil
}

func (us *UserDBService) Login(userData *models.LoginUser) (*models.User, error) {
	for _, user := range us.users {
		// TODO add email hash check
		if user.EmailHash != userData.Email {
			continue
		}

		// TODO add password hash check
		if user.PasswordHash != userData.Password {
			return nil, errors.New("failed to login")
		}

		return user, nil
	}
	return nil, errors.New("failed to login")
}
