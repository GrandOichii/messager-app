package tests

import (
	"testing"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/router"
)

func Test(t *testing.T) {
	r := router.CreateRouter()
	loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	createUser(r, t, "another", "other@mail.com", "pass")

	w, _ := req(r, t, "POST", "/api/users/login", models.PostUser{
		Email:    "mymail@mail.com",
		Password: "1234",
	})

}
