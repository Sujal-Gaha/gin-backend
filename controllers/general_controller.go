package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Gin Backend API!",
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
