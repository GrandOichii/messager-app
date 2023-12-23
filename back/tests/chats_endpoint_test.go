package tests

import (
	"net/http"
	"testing"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/router"
	"github.com/stretchr/testify/assert"
)

func Test_CreateChat(t *testing.T) {
	r := router.CreateRouter()
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	t.Log(token)
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	w, _ := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	})

	t.Logf("%v\n", w.Header())
	assert.Equal(t, http.StatusCreated, w.Code)
}
