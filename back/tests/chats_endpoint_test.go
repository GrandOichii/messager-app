package tests

import (
	"testing"

	"github.com/GrandOichii/messager-app/back/router"
)

func Test(t *testing.T) {
	r := router.CreateRouter()
	loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	createUser(r, t, "another", "other@mail.com", "pass")

}
