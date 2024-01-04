package controllers

import (
	"errors"
	"net/http"

	"github.com/GrandOichii/messager-app/back/connection"
	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

type ChatsController struct {
	Controller

	// UserServicer services.UserServicer
	// ChatServicer services.ChatServicer
	Services *services.Services
	Auth     middleware.Middleware
	Hub      connection.MessageHub
}

func (cs *ChatsController) Map(r *gin.Engine) {
	g := r.Group("/api/chats")

	g.GET("/listen", cs.ListenForMessages)

	gg := g.Group("")
	gg.Use(cs.Auth.GetMiddlewareFunc())
	gg.GET("", cs.GetChatIDs)
	gg.GET(":id", cs.ByID)
	gg.POST("/create", cs.createChat)

	gg.POST("/addmessage", cs.addMessage)

}

func (cs *ChatsController) createChat(c *gin.Context) {
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

	if handle == chatData.WithHandle {
		c.AbortWithError(http.StatusBadRequest, errors.New("User with handle "+handle+" tried to create a chat with themselves"))
		return
	}

	if res, err = cs.Services.ChatServicer.Create(handle, &chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res.ToGetChat())
}

func (cs *ChatsController) addMessage(c *gin.Context) {
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

	cs.Hub.Notify(newMessage.ChatID, res)

	c.JSON(http.StatusCreated, res)
}

func (cs *ChatsController) GetChatIDs(c *gin.Context) {
	handle, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	chatIDs, err := cs.Services.UserServicer.GetChatIDs(handle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, chatIDs)
}

func (cs *ChatsController) ListenForMessages(c *gin.Context) {
	// ? is exposing the chat id like that ok?
	var err error

	chatID := c.Query("chatid")
	if chatID == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("chat id not specified"))
		return
	}

	err = cs.Hub.Register(chatID, c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func (cs *ChatsController) ByID(c *gin.Context) {
	// TODO add message paging to query

	_, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	cID := c.Param("id")
	chat, err := cs.Services.ChatServicer.ByID(cID)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, chat)
}
