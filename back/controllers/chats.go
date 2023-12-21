package controllers

import (
	"net/http"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

type ChatsControllers struct {
	Controller

	UserServicer services.UserServicer
	ChatServicer services.ChatServicer
}

func (cs *ChatsControllers) Map(r *gin.Engine) {
	g := r.Group("/api/chats")

	// TODO require auth
	g.POST("/create", cs.createChat)
	// TODO require auth
	g.POST("/addmessage", cs.addMessage)
}

func (cs *ChatsControllers) createChat(c *gin.Context) {
	var chatData models.CreateChat
	var err error

	if err = c.BindJSON(&chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var res *models.Chat

	// TODO get user id from JWT and use it as owner
	if res, err = cs.ChatServicer.Create(chatData.ByHandle, &chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (cs *ChatsControllers) addMessage(c *gin.Context) {
	var newMessage models.PostMessage

	if err := c.BindJSON(&newMessage); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO get user id from JWT
	user, err := cs.UserServicer.ByHandle(newMessage.OwnerHandle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	chat, err := cs.ChatServicer.ByID(newMessage.ChatID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := cs.ChatServicer.AddMessage(user, chat, &newMessage)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}
