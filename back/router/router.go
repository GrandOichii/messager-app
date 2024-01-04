package router

import (
	"context"
	"time"

	"github.com/GrandOichii/messager-app/back/controllers"
	"github.com/GrandOichii/messager-app/back/middleware"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Router struct {
	Engine   *gin.Engine
	Services *services.Services

	controllers []controllers.Controller

	auth *middleware.JwtMiddleware
}

func CreateRouter() *Router {
	res := gin.Default()

	res.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	result := Router{
		Engine:   res,
		Services: &services.Services{},
	}

	result.configServices(res)
	result.configMiddleware()
	result.createControllers()

	result.configMappings(res)

	return &result
}

func (r *Router) configMiddleware() {
	r.auth = middleware.CreateJwtMiddleware(r.Services)
}

func (r *Router) createControllers() {
	r.controllers = []controllers.Controller{
		&controllers.UsersController{
			Services: r.Services,
			// UserServicer: r.UserServicer,
			Auth: r.auth,
		},
		&controllers.ChatsController{
			Services: r.Services,
			// UserServicer: r.UserServicer,
			// ChatServicer: r.ChatServicer,
			Auth: r.auth,
		},
	}
}

func (r *Router) configMappings(e *gin.Engine) {
	for _, controller := range r.controllers {
		controller.Map(e)
	}

}

func (r *Router) configServices(e *gin.Engine) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// r.Services.UserServicer = services.NewUserService()
	// r.Services.ChatServicer = services.NewChatService(r.Services.UserServicer)
	r.Services.UserServicer = services.NewUserDBService(client)
	r.Services.ChatServicer = services.NewChatDBService(client, r.Services)
}
