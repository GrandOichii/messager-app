package router

import (
	"log"

	"github.com/GrandOichii/messager-app/back/controllers"
	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

var (
	userServicer services.UserServicer
	chatServicer services.ChatServicer

	controllers_ []controllers.Controller
	middleware_  []middleware.Middleware = []middleware.Middleware{
		&middleware.JwtMiddleware{},
	}
)

func CreateRouter() *gin.Engine {
	res := gin.Default()

	configServices(res)
	createControllers()
	configMiddleware(res)
	configMappings(res)

	return res
}

func configMiddleware(r *gin.Engine) {
	for _, mid := range middleware_ {
		err := mid.Apply(r)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func createControllers() {
	controllers_ = []controllers.Controller{
		&controllers.UsersController{
			UserServicer: userServicer,
		},
		&controllers.ChatsControllers{
			UserServicer: userServicer,
			ChatServicer: chatServicer,
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
