package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	Text   string `json:"text"`
	GameID uint   `json:"game_id"`
	UserID uint   `json:"user_id"`
}

// GetAllReviews godoc
// @Summary Get all Reviews.
// @Description Get a list of Reviews.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /Reviews [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Reviews []models.Review
	db.Find(&Reviews)

	c.JSON(http.StatusOK, gin.H{"data": Reviews})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body ReviewInput true "the body to create a new Review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /Reviews [post]
func CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input ReviewInput
	var Reviews models.Game
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&Reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewID not found!"})
		return
	}

	// Create Review
	Review := models.Review{Text: input.Text, GameID: int(input.GameID), UserID: int(input.UserID)}
	db.Create(&Review)

	c.JSON(http.StatusOK, gin.H{"data": Review})
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.Review
// @Router /Reviews/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
	var Review models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Review})
}

//GetReviewByGame godoc
// @Summary  Get Review By Games.
// @Description Get Review by Game.
// @Tags Review
// @Produce json
// @Param id path string true "game id"
// @Success 200 {object} []models.Game
// @Router /Reviews/{id}/Games [get]
func GetReviewByGame(c *gin.Context) { // Get model if exist
	var Games []models.Review

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("game_id = ?", c.Param("id")).Find(&Games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Games})
}

//GetReviewByUser godoc
// @Summary  Get Review By User.
// @Description Get Review by User.
// @Tags Review
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} []models.User
// @Router /Reviews/{id}/Users [get]
func GetReviewByUser(c *gin.Context) { // Get model if exist
	var Users []models.Review

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", c.Param("id")).Find(&Users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Users})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Param Body body ReviewInput true "the body to update an Review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Review
// @Router /Reviews/{id} [patch]
func UpdateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Review models.Review
	var Reviews models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&Review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&Reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewID not found!"})
		return
	}

	var updatedInput models.Review
	updatedInput.Text = input.Text
	updatedInput.GameID = int(input.GameID)
	updatedInput.UserID = int(input.UserID)
	updatedInput.UpdatedAt = time.Now()

	db.Model(&Review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Review})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a Review by id.
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} map[string]boolean
// @Router /Review/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&Review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
