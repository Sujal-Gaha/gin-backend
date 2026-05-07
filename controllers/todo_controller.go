package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gin-backend/models"
)

var (
	todos       = []models.Todo{}
	nextTodoID  = 1
)

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = nextTodoID
	nextTodoID++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			updatedTodo.ID = id
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
