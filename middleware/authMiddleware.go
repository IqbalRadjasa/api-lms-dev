package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Read token from cookie
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authentication required",
			})
			return
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
			})
			return
		}

		// 3. Extract claims safely
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token claims",
			})
			return
		}

		// 4. Convert claims
		userID, ok1 := claims["user_id"].(float64)
		roleID, ok2 := claims["role_id"].(float64)

		if !ok1 || !ok2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token payload",
			})
			return
		}

		// 5. Store values in context
		c.Set("user_id", int(userID))
		c.Set("role_id", int(roleID))

		c.Next()
	}
}


func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("role_id")
		if !exists || roleID.(int) != 1 {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied. Admin only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
