package controllers

import (
	"api-lms-dev/database"
	"api-lms-dev/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// == Roles ==
func GetAllRoles(c *gin.Context) {
	var roles []models.Roles

	err := database.DB.Find(&roles).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No datas found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func GetRoleById(c *gin.Context) {
	id := c.Param("id")

	var role models.Roles
	err := database.DB.First(&role, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Role not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func CreateRole(c *gin.Context) {
	validate := validator.New()

	var req models.Roles
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// Insert data
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create role",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Role created successfully",
		"data":    req,
	})
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")

	var req models.Roles
	if err := database.DB.Where("id = ?", id).First(&req).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Role not found"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	delete(updateData, "id")
	delete(updateData, "created_at")
	delete(updateData, "updated_at")

	database.DB.Model(&req).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
		"data":    req,
	})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Roles{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}

// == Departments ==
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

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	var dept models.Departments
	err := database.DB.First(&dept, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Department not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dept})
}

func CreateDepartment(c *gin.Context) {
	validate := validator.New()

	var req models.Departments
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// Insert data
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create department",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Department created successfully",
		"data":    req,
	})
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")

	var req models.Departments
	if err := database.DB.Where("id = ?", id).First(&req).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Department not found"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	delete(updateData, "id")
	delete(updateData, "created_at")
	delete(updateData, "updated_at")

	database.DB.Model(&req).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
		"data":    req,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Departments{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted"})
}

// == Categories ==
func GetAllCategories(c *gin.Context) {
	var categories []models.Categories

	err := database.DB.Preload("Department").Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No datas found!"})
		return
	}

	type responseDTO struct {
		ID         int    `json:"id"`
		Department string `json:"department"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
	}

	var response []responseDTO
	for _, c := range categories {
		response = append(response, responseDTO{
			ID:         c.ID,
			Department: c.Department.Nickname,
			Name:       c.Name,
			Slug:       c.Slug,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	var category models.Categories
	err := database.DB.Preload("Department").First(&category, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found!"})
		return
	}

	type responseDTO struct {
		ID         int    `json:"id"`
		Department string `json:"department"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
	}

	response := responseDTO{
		ID:         category.ID,
		Department: category.Department.Nickname,
		Name:       category.Name,
		Slug:       category.Slug,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func CreateCategory(c *gin.Context) {
	validate := validator.New()

	var req models.Categories
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// Insert data
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create category",
			"error":   err.Error(),
		})
		return
	}
	type responseDTO struct {
		ID           int    `json:"id"`
		DepartmentId int    `json:"department_id"`
		Name         string `json:"name"`
		Slug         string `json:"slug"`
	}

	response := responseDTO{
		ID:           req.ID,
		DepartmentId: req.DepartmentId,
		Name:         req.Name,
		Slug:         req.Slug,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    response,
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var req models.Categories
	if err := database.DB.Where("id = ?", id).First(&req).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	delete(updateData, "id")
	delete(updateData, "created_at")
	delete(updateData, "updated_at")

	database.DB.Model(&req).Updates(updateData)

	type responseDTO struct {
		ID           int    `json:"id"`
		DepartmentId int    `json:"department_id"`
		Name         string `json:"name"`
		Slug         string `json:"slug"`
	}

	response := responseDTO{
		ID:           req.ID,
		DepartmentId: req.DepartmentId,
		Name:         req.Name,
		Slug:         req.Slug,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
		"data":    response,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Categories{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
