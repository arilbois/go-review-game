package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TopReviewInput struct {
	Name     string `json:"name"`
	ReviewID uint   `json:"review_ID"`
}

// GetAllTopReview godoc
// @Summary Get all TopReview.
// @Description Get a list of TopReview.
// @Tags TopReview
// @Produce json
// @Success 200 {object} []models.TopReview
// @Router /TopReview [get]
func GetAllTopR(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var TopR []models.TopReview
	db.Find(&TopR)

	c.JSON(http.StatusOK, gin.H{"data": TopR})
}

// CreateTopReview godoc
// @Summary Create New TopReview.
// @Description Creating a new TopReview.
// @Tags TopReview
// @Param Body body TopReviewInput true "the body to create a new TopReview"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.TopReview
// @Router /TopReview [post]
func CreateTopR(c *gin.Context) {
	// Validate input
	var input TopReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create TopR
	TopR := models.TopReview{Name: input.Name, ReviewID: input.ReviewID}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&TopR)

	c.JSON(http.StatusOK, gin.H{"data": TopR})
}

// UpdateTopReview godoc
// @Summary Update TopReview.
// @Description Update TopReview by id.
// @Tags TopReview
// @Produce json
// @Param id path string true "TopReview id"
// @Param Body body TopReviewInput true "the body to update age TopR TopReview"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.TopReview
// @Router /TopReview/{id} [patch]
func UpdateTopR(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var TopR models.TopReview
	if err := db.Where("id = ?", c.Param("id")).First(&TopR).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input TopReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.TopReview
	updatedInput.Name = input.Name
	updatedInput.ReviewID = input.ReviewID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&TopR).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": TopR})
}

// DeleteTopReview godoc
// @Summary Delete one TopReview.
// @Description Delete a TopReview by id.
// @Tags TopReview
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "TopReview id"
// @Success 200 {object} map[string]boolean
// @Router /TopReview/{id} [delete]
func DeleteTopR(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var TopR models.TopReview
	if err := db.Where("id = ?", c.Param("id")).First(&TopR).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&TopR)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
