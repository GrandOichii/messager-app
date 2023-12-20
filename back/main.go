package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []User

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func getUsers(c *gin.Context) {
}
