package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/dto"
	"api-lms-dev/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GET all users
func GetAllUsers(c *gin.Context) {
	var users []models.Users
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// POST create new user
func CreateUser(c *gin.Context) {
	type RequestBody struct {
		Users       models.Users       `json:"users"`
		UserDetails models.UserDetails `json:"user_details"`
	}

	var req RequestBody
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Users.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	req.Users.Password = string(hashedPassword)

	// Begin transaction
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		// Insert user
		if err := tx.Create(&req.Users).Error; err != nil {
			return err
		}

		// Set foreign key
		req.UserDetails.UserId = req.Users.ID

		// Insert user details
		if err := tx.Create(&req.UserDetails).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user",
			"error":   err.Error(),
		})
		return
	}

	userResponse := dto.UserResponse{
		ID:     req.Users.ID,
		Nisn:   req.Users.Nisn,
		RoleId: req.Users.RoleId,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "User created successfully",
		"user":         userResponse,
		"user_details": req.UserDetails,
	})
}

// GET user by id
func GetUserById(c *gin.Context) {
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
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	database.DB.First(&user, id)

	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DELETE user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Users{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
