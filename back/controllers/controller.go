package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Map(r *gin.Engine)
}
