package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetAllCategory godoc
// @Summary Get all Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /Catagories [get]
func GetAllCat(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Cat []models.Category
	db.Find(&Cat)

	c.JSON(http.StatusOK, gin.H{"data": Cat})
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body CategoryInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /Catagories [post]
func CreateCat(c *gin.Context) {
	// Validate input
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Cat
	Cat := models.Category{Name: input.Name, Description: input.Description}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Cat)

	c.JSON(http.StatusOK, gin.H{"data": Cat})
}

// GetCategoryById godoc
// @Summary Get Category.
// @Description Get an Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /Catagories/{id} [get]
func GetCatById(c *gin.Context) { // Get model if exist
	var Cat models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Cat})
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param Body body CategoryInput true "the body to update age Cat category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Category
// @Router /Catagories/{id} [patch]
func UpdateCat(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Cat models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&Cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.UpdatedAt = time.Now()

	db.Model(&Cat).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Cat})
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} map[string]boolean
// @Router /Catagories/{id} [delete]
func DeleteCat(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Cat models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&Cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Cat)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
