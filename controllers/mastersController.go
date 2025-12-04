package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Roles
func GetAllRoles(c *gin.Context) {
	var roles []models.Roles

	err := database.DB.Find(&roles).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No datas found!"})
		return
	}

	type responseDTO struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var response []responseDTO
	for _, role := range roles {
		response = append(response, responseDTO{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
