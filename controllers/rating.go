package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RatingInput struct {
	Value string `json:"value"`
}

// GetAllRating godoc
// @Summary Get all Rating.
// @Description Get a list of Rating.
// @Tags Rating
// @Produce json
// @Success 200 {object} []models.Rating
// @Router /Rating [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Rating []models.Rating
	db.Find(&Rating)

	c.JSON(http.StatusOK, gin.H{"data": Rating})
}

// CreateRating godoc
// @Summary Create New Rating.
// @Description Creating a new Rating.
// @Tags Rating
// @Param Body body RatingInput true "the body to create a new Rating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Rating
// @Router /Rating [post]
func CreateRating(c *gin.Context) {
	// Validate input
	var input RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	Rating := models.Rating{Value: input.Value}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Rating)

	c.JSON(http.StatusOK, gin.H{"data": Rating})
}

// UpdateRating godoc
// @Summary Update Rating.
// @Description Update Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Param Body body RatingInput true "the body to update age Rating Rating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Rating
// @Router /Rating/{id} [patch]
func UpdateRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&Rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Rating
	updatedInput.Value = input.Value
	updatedInput.UpdatedAt = time.Now()

	db.Model(&Rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Rating})
}

// DeleteRating godoc
// @Summary Delete one Rating.
// @Description Delete a Rating by id.
// @Tags Rating
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} map[string]boolean
// @Router /Rating/{id} [delete]
func DeleteRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&Rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
