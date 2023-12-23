package services

import "github.com/GrandOichii/messager-app/back/models"

type UserServicer interface {
	// TODO remove
	All() []*models.GetUser
	ByHandle(handle string) (*models.User, error)

	Register(newUser *models.CreateUser) (*models.GetUser, error)
	Login(user *models.LoginUser) (*models.User, error)
}
