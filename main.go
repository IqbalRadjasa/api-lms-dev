package main

import (
	"api-lms-dev/controllers"
	"api-lms-dev/database"
	"api-lms-dev/models"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.Users{})

	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserById)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":8085")
}