package router

import (
	"github.com/GrandOichii/messager-app/back/controllers"
	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

var (
	userServicer services.UserServicer
	chatServicer services.ChatServicer

	controllers_ []controllers.Controller

	auth *middleware.JwtMiddleware
)

func CreateRouter() *gin.Engine {
	res := gin.Default()

	configServices(res)
	configMiddleware()
	createControllers()

	configMappings(res)

	return res
}

func configMiddleware() {
	auth = middleware.CreateJwtMiddleware(userServicer)
}

func createControllers() {
	controllers_ = []controllers.Controller{
		&controllers.UsersController{
			UserServicer: userServicer,
			Auth:         auth,
		},
		&controllers.ChatsControllers{
			UserServicer: userServicer,
			ChatServicer: chatServicer,
			Auth:         auth,
		},
	}
}

func configMappings(r *gin.Engine) {
	for _, controller := range controllers_ {
		controller.Map(r)
	}

}

func configServices(r *gin.Engine) {
	userServicer = services.NewUserService()
	chatServicer = services.NewChatService()
}
