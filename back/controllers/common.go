package controllers

import (
	"errors"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func extract(key string, c *gin.Context) (string, error) {
	claims := jwt.ExtractClaims(c)
	result := claims[key]
	if result == nil {
		return "", errors.New("key " + key + " doesn't exist")
	}

	return result.(string), nil
}
