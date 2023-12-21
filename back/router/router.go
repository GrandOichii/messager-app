package router

import (
	"net/http"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

var (
	userServicer services.UserServicer
	chatServicer services.ChatServicer
)

func CreateRouter() *gin.Engine {
	res := gin.Default()

	configMappings(res)
	configServices(res)

	return res
}

func configMappings(r *gin.Engine) {

	// TODO require auth
	r.GET("/api/users", getUsers)

	r.POST("/api/users/register", registerUser)
	r.POST("/api/users/login", loginUser)

	// TODO require auth
	r.POST("/api/chats/create", createChat)
	// TODO require auth
	r.POST("/api/chats/addmessage", addMessage)
}

func configServices(r *gin.Engine) {
	userServicer = services.NewUserService()
	chatServicer = services.NewChatService()
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userServicer.All())
}

func registerUser(c *gin.Context) {
	var userData models.CreateUser
	var err error

	if err = c.BindJSON(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newUser *models.GetUser
	if newUser, err = userServicer.Register(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func loginUser(c *gin.Context) {
	var userData models.PostUser
	var err error

	if err = c.BindJSON(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var token string
	if token, err = userServicer.Login(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func createChat(c *gin.Context) {
	var chatData models.CreateChat
	var err error

	if err = c.BindJSON(&chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var res *models.Chat

	// TODO get user id from JWT and use it as owner
	if res, err = chatServicer.Create(chatData.ByHandle, &chatData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func addMessage(c *gin.Context) {
	var newMessage models.PostMessage

	if err := c.BindJSON(&newMessage); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO get user id from JWT
	user, err := userServicer.ByHandle(newMessage.OwnerHandle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	chat, err := chatServicer.ByID(newMessage.ChatID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := chatServicer.AddMessage(user, chat, &newMessage)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}
