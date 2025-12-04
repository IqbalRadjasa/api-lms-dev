package main

import (
	"log"
	"os"

	"api-lms-dev/controllers"
	"api-lms-dev/database"
	"api-lms-dev/middleware"
	"api-lms-dev/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.Users{})

	// Public Route
	r.POST("/login", controllers.Login)

	// Protected Routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// User Management
	authorized.GET("/users", middleware.AdminOnly(), controllers.GetAllUsers)
	authorized.GET("/users/:id", controllers.GetUserById)
	authorized.POST("/users", middleware.AdminOnly(), controllers.CreateUser)
	authorized.PUT("/users/:id", controllers.UpdateUserDetail)
	authorized.DELETE("/users/:id", middleware.AdminOnly(), controllers.DeleteUser)

	// Masters
	authorized.GET("/master/roles", controllers.GetAllRoles)
	authorized.GET("/master/departments", controllers.GetAllDepartments)
	authorized.GET("/master/categories", controllers.GetAllCategories)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPort := os.Getenv("API_PORT")
	r.Run(":" + apiPort)
}
