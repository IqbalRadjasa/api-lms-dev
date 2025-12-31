package main

import (
	"log"
	"os"

	"api-lms-dev/controllers"
	"api-lms-dev/database"
	"api-lms-dev/middleware"
	"api-lms-dev/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Users{})

	// r := gin.Default()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))


	// Public Route
	app.Post("/login", controllers.Login)

	// == Protected Routes ==
	authorized := app.Group("/")
	authorized.Use(middleware.AuthMiddleware)

	authorized.Get("/auth/validate", controllers.Validate)
	authorized.Get("/auth/me", controllers.Me)

	// // User Management
	// authorized.GET("/users", middleware.AdminOnly(), controllers.GetAllUsers)
	// authorized.GET("/users/:id", controllers.GetUserById)
	// authorized.POST("/users", middleware.AdminOnly(), controllers.CreateUser)
	// authorized.PUT("/users/:id", controllers.UpdateUserDetail)
	// authorized.DELETE("/users/:id", middleware.AdminOnly(), controllers.DeleteUser)

	// Logout
	authorized.Post("/logout", controllers.Logout)

	// // == Course ==
	// authorized.POST("/courses", controllers.CreateCourse)

	// // == Masters ==
	// // Roles
	// authorized.GET("/master/roles", controllers.GetAllRoles)
	// authorized.GET("/master/roles/:id", controllers.GetRoleById)
	// authorized.POST("/master/roles", controllers.CreateRole)
	// authorized.DELETE("/master/roles/:id", controllers.DeleteRole)
	// authorized.PUT("/master/roles/:id", controllers.UpdateRole)

	// // Departments
	// authorized.GET("/master/departments", controllers.GetAllDepartments)
	// authorized.GET("/master/departments/:id", controllers.GetDepartmentById)
	// authorized.POST("/master/departments", controllers.CreateDepartment)
	// authorized.PUT("/master/departments/:id", controllers.UpdateDepartment)
	// authorized.DELETE("/master/departments/:id", controllers.DeleteDepartment)

	// // Categories
	// authorized.GET("/master/categories", controllers.GetAllCategories)
	// authorized.GET("/master/categories/:id", controllers.GetCategoryById)
	// authorized.POST("/master/categories", controllers.CreateCategory)
	// authorized.PUT("/master/categories/:id", controllers.UpdateCategory)
	// authorized.DELETE("/master/categories/:id", controllers.DeleteCategory)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPort := os.Getenv("API_PORT")

	log.Println("Server running on port", apiPort)
	app.Listen(":" + apiPort)
}
