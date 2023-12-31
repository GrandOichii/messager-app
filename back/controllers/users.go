package controllers

import (
	"net/http"

	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	Controller

	// UserServicer services.UserServicer
	Services *services.Services
	Auth     *middleware.JwtMiddleware
}

func (uc *UsersController) Map(r *gin.Engine) {
	g := r.Group("/api/users")

	g.GET("", uc.getUsers)

	g.POST("/register", uc.registerUser)
	g.POST("/login", uc.Auth.Middle.LoginHandler)

	g.GET("/avatar/:uhandle", uc.GetAvatar)
}

func (uc *UsersController) getUsers(c *gin.Context) {
	// _, err := extract(middleware.IDKey, c)
	// if err != nil {
	// 	c.AbortWithError(http.StatusUnauthorized, err)
	// 	return
	// }

	users, err := uc.Services.UserServicer.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UsersController) GetAvatar(c *gin.Context) {
	handle := c.Param("uhandle")
	user, err := uc.Services.UserServicer.ByHandle(handle)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.String(http.StatusOK, user.AvatarURI)
}

func (uc *UsersController) registerUser(c *gin.Context) {
	var userData models.CreateUser
	var err error

	if err = c.BindJSON(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = userData.CheckValid()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newUser *models.GetUser
	if newUser, err = uc.Services.UserServicer.Register(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
