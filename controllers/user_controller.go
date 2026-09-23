package controllers

import (
	"net/http"
	"strconv"

	"gin-backend/models"
	"github.com/gin-gonic/gin"
)

var (
	users  = []models.User{}
	nextID = 1
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
