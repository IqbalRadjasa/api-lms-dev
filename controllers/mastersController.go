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
	for _, r := range roles {
		response = append(response, responseDTO{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// Get All Departments
func GetAllDepartments(c *gin.Context) {
	var departments []models.Departments

	err := database.DB.Find(&departments).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No datas found!"})
		return
	}

	type responseDTO struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Nickname string `json:"nickname"`
		Slug     string `json:"slug"`
	}

	var response []responseDTO
	for _, d := range departments {
		response = append(response, responseDTO{
			ID:       d.ID,
			Name:     d.Name,
			Nickname: d.Nickname,
			Slug:     d.Slug,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
