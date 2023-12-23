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

	UserServicer services.UserServicer
	Auth         middleware.Middleware
}

func (uc *UsersController) Map(r *gin.Engine) {
	g := r.Group("/api/users")

	g.GET("", uc.getUsers)

	g.POST("/register", uc.registerUser)
	g.POST("/login", uc.loginUser)
}

func (uc *UsersController) getUsers(c *gin.Context) {
	_, err := extract(middleware.IDKey, c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, uc.UserServicer.All())
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
	if newUser, err = uc.UserServicer.Register(&userData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (uc *UsersController) loginUser(c *gin.Context) {
	var userData models.PostUser
	var err error

	if err = c.BindJSON(&userData); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = userData.CheckValid()
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var token string
	if token, err = uc.UserServicer.Login(&userData); err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
