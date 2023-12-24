package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/stretchr/testify/assert"
)

func Test_CreateChat(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")
	w, _ := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_CreateChat_Failed(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	w, _ := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: "non-existant",
	}, token)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_CreateChat_AlreadyExists(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)
	w, _ := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_SendMessage(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	_, chatData := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)

	var chat models.Chat
	err := json.Unmarshal(chatData, &chat)
	checkErr(t, err)

	w, _ := req(r, t, "POST", "/api/chats/addmessage", models.PostMessage{
		ChatID: chat.ID,
		Text:   "Hello, world!",
	}, token)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_SendMessage_NoChatId(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	_, chatData := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)

	var chat models.Chat
	err := json.Unmarshal(chatData, &chat)
	checkErr(t, err)

	w, _ := req(r, t, "POST", "/api/chats/addmessage", models.PostMessage{
		ChatID: "invalid-id",
		Text:   "Hello, world!",
	}, token)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_SendMessage_NoMessage(t *testing.T) {
	r := createRouter().Engine
	token := loginAs(r, t, "coolhandle", "mymail@mail.com", "1234")
	otherHandle := "another"
	createUser(r, t, otherHandle, "other@mail.com", "pass")

	_, chatData := req(r, t, "POST", "/api/chats/create", models.CreateChat{
		WithHandle: otherHandle,
	}, token)

	var chat models.Chat
	err := json.Unmarshal(chatData, &chat)
	checkErr(t, err)

	w, _ := req(r, t, "POST", "/api/chats/addmessage", models.PostMessage{
		ChatID: chat.ID,
		Text:   "",
	}, token)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
