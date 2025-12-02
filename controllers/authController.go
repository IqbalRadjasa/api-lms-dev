package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"api-lms-dev/database"
	"api-lms-dev/models"
)

func Login(c *gin.Context) {
	type LoginRequest struct {
		Nisn_or_Nip string `json:"nisn_or_nip"`
		Password    string `json:"password"`
	}

	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Find user by NISN OR NIP
	var user models.Users
	if err := database.DB.Where("nisn = ? OR nip = ?", req.Nisn_or_Nip, req.Nisn_or_Nip).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role_id": user.RoleId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   tokenString,
		// "user":    user,
	})
}
