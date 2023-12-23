package services

import (
	"errors"

	"github.com/GrandOichii/messager-app/back/models"
)

type UserService struct {
	UserServicer

	users []*models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []*models.User{
			{
				Handle:       "handle1",
				PasswordHash: "pass",
				EmailHash:    "mail@mail.com",
			},
		},
	}
}

func (us *UserService) All() []*models.GetUser {
	res := make([]*models.GetUser, len(us.users))
	for i, user := range us.users {
		res[i] = user.ToGetUser()
	}
	return res
}

func (us *UserService) ByHandle(handle string) (*models.User, error) {
	for _, user := range us.users {
		if user.Handle == handle {
			return user, nil
		}
	}
	return nil, errors.New("no user with handle " + handle)
}

func (us *UserService) Register(newUser *models.CreateUser) (*models.GetUser, error) {
	for _, u := range us.users {
		if u.Handle == newUser.Handle {
			return nil, errors.New("User with handle " + newUser.Handle + " already exists")
		}
		// TODO chack mail hash
		if u.EmailHash == newUser.Email {
			return nil, errors.New("User with handle " + newUser.Handle + " already exists")
		}
	}

	res := &models.User{
		Handle: newUser.Handle,
		// TODO hash email
		EmailHash: newUser.Email,
		// TODO hash password
		PasswordHash: newUser.Password,
	}

	us.users = append(us.users, res)
	return res.ToGetUser(), nil
}

func (us *UserService) Login(userData *models.LoginUser) (*models.User, error) {
	// fmt.Printf("userData.Email: %v\n", userData.Email)
	for _, user := range us.users {
		// fmt.Printf("\tuser.EmailHash: %v\n", user.EmailHash)
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
