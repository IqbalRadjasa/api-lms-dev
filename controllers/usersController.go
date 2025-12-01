package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET all users
func GetAllUsers(c *gin.Context){
	var users []models.Users
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// POST create new user
func CreateUser(c *gin.Context){
	var user models.Users
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// GET user by id
func GetUserById(c *gin.Context){
	id := c.Param("id")
	var user models.Users
	result := database.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// PUT user
func UpdateUser(c *gin.Context){
	id := c.Param("id")
	var user models.Users
	database.DB.First(&user, id)

	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DELETE user
func DeleteUser(c *gin.Context){
	id := c.Param("id")
	database.DB.Delete(&models.Users{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}