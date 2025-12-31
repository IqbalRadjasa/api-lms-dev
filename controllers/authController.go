package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"api-lms-dev/database"
	"api-lms-dev/models"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	var req LoginRequest

	// Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	// Find user by NISN or NIP
	var user models.Users
	if err := database.DB.Where("identifier = ?", req.Identifier).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "NISN / NIP tidak ditemukan",
		})
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		 "role_id": user.RoleId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   86400, // 24 hours
		HTTPOnly: true,
		Secure:   false, // set true if using HTTPS
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"token":   tokenString,
	})
}

func Validate(c *fiber.Ctx) error {
	userId := c.Locals("user_id")

	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Authorized",
	})
}

func Me(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		});
	}

	convertedUserId := userId.(int)

	var user models.Users
	err := database.DB.
		Preload("Role").
		Preload("UserDetails").
		Preload("UserDetails.Department").
		First(&user, convertedUserId).
		Error

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	data := fiber.Map{
		"id":         user.ID,
		"identifier": user.Identifier,
		"role": fiber.Map{
			"id":   user.Role.ID,
			"name": user.Role.Name,
		},
		"user_details": fiber.Map{
			"fullname":   user.UserDetails.Fullname,
			"nickname":   user.UserDetails.Nickname,
			"department": user.UserDetails.Department.Name,
			"email":      user.UserDetails.Email,
			"phone":      user.UserDetails.Phone,
		},
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("access_token")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout success",
	})
}


