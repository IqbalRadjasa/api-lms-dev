package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/models"
	"path"
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCourse(c *gin.Context) {
	// Parse form
	categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
	instructorId, _ := strconv.Atoi(c.PostForm("instructor_id"))
	slug := c.PostForm("slug")
	title := c.PostForm("title")
	description := c.PostForm("description")

	// Handle file upload
	file, err := c.FormFile("thumbnail")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thumbnail is required"})
		return
	}

	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	// Validate extension
	ext := strings.ToLower(path.Ext(file.Filename))
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExt[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file type. Allowed types: .jpg .jpeg .png .webp",
		})
		return
	}

	// Save file to local folder /uploads
	filename := "uploads/thumbnails/" + uuid.New().String() + path.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Create model instance
	course := models.Courses{
		CategoryId:   categoryId,
		InstructorId: instructorId,
		Slug:         slug,
		Title:        title,
		Description:  description,
		Thumbnail:    filename, // store path in DB
	}

	// Insert data
	if err := database.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create course",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
		"data":    course,
	})
}
