package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {

	// 1. Read token from cookie
	tokenString := c.Cookies("access_token")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "authentication required",
		})
	}

	secret := os.Getenv("JWT_SECRET")

	// 2. Parse & validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	// 3. Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		 return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token claims",
		})
	}

	// 4. Convert values
	userID, ok1 := claims["user_id"].(float64)
	roleID, ok2 := claims["role_id"].(float64)

	if !ok1 || !ok2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token payload",
		})
	}

	// 5. Store in Fiber context
	c.Locals("user_id", int(userID))
	c.Locals("role_id", int(roleID))

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {

	roleID := c.Locals("role_id")

	if roleID == nil || roleID.(int) != 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied. Admin only",
		})
	}

	return c.Next()
}
