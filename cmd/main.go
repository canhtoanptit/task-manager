package main

import (
	"task-management-api/internal/config"
	"task-management-api/internal/database"
	"task-management-api/internal/handlers"
	"task-management-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	r := gin.Default()

	// Public routes
	r.POST("/login", handlers.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.Authenticate)
	{
		api.GET("/tasks", handlers.GetTasks)
		api.GET("/tasks/my-tasks", handlers.GetMyTasks)
		api.POST("/tasks", middleware.EmployerRequired, handlers.CreateTask)
		api.PATCH("/tasks/:id", middleware.EmployeeRequired, handlers.UpdateTask)
		api.GET("/tasks/summary", middleware.EmployerRequired, handlers.GetEmployeeSummary)
	}

	r.Run(":8000")
}
