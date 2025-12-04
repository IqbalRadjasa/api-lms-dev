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

	// == Protected Routes ==
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// User Management
	authorized.GET("/users", middleware.AdminOnly(), controllers.GetAllUsers)
	authorized.GET("/users/:id", controllers.GetUserById)
	authorized.POST("/users", middleware.AdminOnly(), controllers.CreateUser)
	authorized.PUT("/users/:id", controllers.UpdateUserDetail)
	authorized.DELETE("/users/:id", middleware.AdminOnly(), controllers.DeleteUser)

	// == Course ==
	authorized.POST("/courses", controllers.CreateCourse)

	// == Masters ==
	// Roles
	authorized.GET("/master/roles", controllers.GetAllRoles)
	authorized.GET("/master/roles/:id", controllers.GetRoleById)
	authorized.POST("/master/roles", controllers.CreateRole)
	authorized.PUT("/master/roles/:id", controllers.UpdateRole)
	authorized.DELETE("/master/roles/:id", controllers.DeleteRole)

	// Departments
	authorized.GET("/master/departments", controllers.GetAllDepartments)
	authorized.GET("/master/departments/:id", controllers.GetDepartmentById)
	authorized.POST("/master/departments", controllers.CreateDepartment)
	authorized.PUT("/master/departments/:id", controllers.UpdateDepartment)
	authorized.DELETE("/master/departments/:id", controllers.DeleteDepartment)

	// Categories
	authorized.GET("/master/categories", controllers.GetAllCategories)
	authorized.GET("/master/categories/:id", controllers.GetCategoryById)
	authorized.POST("/master/categories", controllers.CreateCategory)
	authorized.PUT("/master/categories/:id", controllers.UpdateCategory)
	authorized.DELETE("/master/categories/:id", controllers.DeleteCategory)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPort := os.Getenv("API_PORT")
	r.Run(":" + apiPort)
}
