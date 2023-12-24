package router

import (
	"context"

	"github.com/GrandOichii/messager-app/back/controllers"
	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// userServicer = services.NewUserService()
	userServicer = services.NewUserDBService(client)
	chatServicer = services.NewChatService(userServicer)
}
