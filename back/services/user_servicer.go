package services

import "github.com/GrandOichii/messager-app/back/models"

type UserServicer interface {
	// TODO remove
	All() []*models.GetUser

	Register(newUser *models.CreateUser) (*models.GetUser, error)
	Login(user *models.PostUser) (string, error)
}
