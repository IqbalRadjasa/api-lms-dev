package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/dto"
	"api-lms-dev/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()

	type RequestBody struct {
		Users       models.Users       `json:"users" validate:"required"`
		UserDetails models.UserDetails `json:"user_details" validate:"required"`
	}

	var req RequestBody
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// Hashing password
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

	var nisnOrNip string
	if req.Users.RoleId == 1 || req.Users.RoleId == 2 {
		nisnOrNip = req.Users.Nip
	} else {
		nisnOrNip = req.Users.Nisn
	}

	userResponse := dto.UserResponse{
		ID:          req.Users.ID,
		Nisn_or_Nip: nisnOrNip,
		RoleId:      req.Users.RoleId,
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
	if err := database.DB.Preload("UserDetails").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
		return
	}

	var nisnOrNip string
	if user.RoleId == 1 || user.RoleId == 2 {
		nisnOrNip = user.Nip
	} else {
		nisnOrNip = user.Nisn
	}

	response := dto.UserDetailResponse{
		ID:           user.UserDetails.ID,
		UserId:       user.UserDetails.UserId,
		NisnOrNip:    nisnOrNip,
		Fullname:     user.UserDetails.Fullname,
		Nickname:     user.UserDetails.Nickname,
		DateOfBirth:  user.UserDetails.DateOfBirth,
		PlaceOfBirth: user.UserDetails.PlaceOfBirth,
		Email:        user.UserDetails.Email,
		Phone:        user.UserDetails.Phone,
		Address:      user.UserDetails.Address,
	}

	c.JSON(http.StatusOK, response)
}

// PUT user
func UpdateUserDetail(c *gin.Context) {
	id := c.Param("id")

	var userDetail models.UserDetails
	if err := database.DB.Where("user_id = ?", id).First(&userDetail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User detail not found"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	delete(updateData, "user_id")
	delete(updateData, "id")
	delete(updateData, "created_at")
	delete(updateData, "updated_at")

	database.DB.Model(&userDetail).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
		"data":    userDetail,
	})
}

// DELETE user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Users{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
