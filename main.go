package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sujal/gin-backend-v1/controllers"
)

func main() {
	r := gin.Default()

	// General routes
	r.GET("/", controllers.Greet)
	r.GET("/health", controllers.HealthCheck)

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	// Todo routes
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetTodo)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}

	r.Run(":8080")
}
