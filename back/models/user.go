package models

// TODO move all DTOs to a separate folder

type User struct {
	Handle       string   `json:"handle"`
	PasswordHash string   `json:"passhash"`
	EmailHash    string   `json:"emailhash"`
	ChatIDs      []string `json:"chats_ids"`
}

func (u *User) ToGetUser() *GetUser {
	return &GetUser{
		Handle: u.Handle,
	}
}

type GetUser struct {
	Handle string `json:"handle"`
}

type PostUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password"`
	Handle   string `json:"handle"`
}
