package controllers

import (
	"net/http"

	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

type ChatsControllers struct {
	Controller

	// UserServicer services.UserServicer
	// ChatServicer services.ChatServicer
	Services *services.Services
	Auth     middleware.Middleware
}

func (cs *ChatsControllers) Map(r *gin.Engine) {
	g := r.Group("/api/chats")

	g.Use(cs.Auth.GetMiddlewareFunc())
	g.GET("", cs.GetChatIDs)
	g.POST("/create", cs.createChat)

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

	handle, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if res, err = cs.Services.ChatServicer.Create(handle, &chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res.ToGetChat())
	// c.JSON(http.StatusCreated, gin.H{})
}

func (cs *ChatsControllers) addMessage(c *gin.Context) {
	var newMessage models.PostMessage

	if err := c.BindJSON(&newMessage); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	handle, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	user, err := cs.Services.UserServicer.ByHandle(handle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	chat, err := cs.Services.ChatServicer.ByID(newMessage.ChatID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := cs.Services.ChatServicer.AddMessage(user, chat, &newMessage)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (cs *ChatsControllers) GetChatIDs(c *gin.Context) {
	handle, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	chatIDs, err := cs.Services.UserServicer.GetChatIDs(handle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, chatIDs)
}
