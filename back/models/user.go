package models

import "errors"

// TODO move all DTOs to a separate folder

type User struct {
	AvatarURI    string   `json:"avatar_uri" bson:"avatar_uri"`
	Handle       string   `json:"handle" bson:"handle"`
	PasswordHash string   `json:"passhash" bson:"passhash"`
	EmailHash    string   `json:"emailhash" bson:"emailhash"`
	ChatIDs      []string `json:"chat_ids" bson:"chat_ids"`
}

func (u *User) ToGetUser() *GetUser {
	return &GetUser{
		Handle: u.Handle,
	}
}

type GetUser struct {
	Handle string `json:"handle"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *LoginUser) CheckValid() error {
	if len(u.Email) == 0 {
		return errors.New("invalid email")
	}
	if len(u.Password) == 0 {
		return errors.New("invalid password")
	}
	return nil
}

type CreateUser struct {
	// TODO takes up too much space, move to separate file
	AvatarURI string `json:"avatar_uri" bson:"avatar_uri"`
	Email     string `json:"email" required:"true"`
	Password  string `json:"password"`
	Handle    string `json:"handle"`
}

func (u *CreateUser) CheckValid() error {
	if len(u.AvatarURI) == 0 {
		return errors.New("can't create user with no avatar")
	}
	if len(u.Email) == 0 {
		return errors.New("can't create user with no email")
	}
	if len(u.Password) == 0 {
		return errors.New("can't create user with no password")
	}
	if len(u.Handle) == 0 {
		return errors.New("can't create user with no handle")
	}
	return nil
}
