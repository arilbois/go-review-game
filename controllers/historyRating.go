package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HistoryRatingInput struct {
	RatingID uint `json:"rating_id"`
	GameID   uint `json:"game_id"`
	UserID   uint `json:"user_id"`
}

// GetAllHistoryRatings godoc
// @Summary Get all HistoryRatings.
// @Description Get a list of HistoryRatings.
// @Tags HistoryRating
// @Produce json
// @Success 200 {object} []models.HistoryRating
// @Router /HistoryRatings [get]
func GetAllHistoryRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var HistoryRatings []models.HistoryRating
	db.Find(&HistoryRatings)

	c.JSON(http.StatusOK, gin.H{"data": HistoryRatings})
}

// CreateHistoryRating godoc
// @Summary Create New HistoryRating.
// @Description Creating a new HistoryRating.
// @Tags HistoryRating
// @Param Body body HistoryRatingInput true "the body to create a new HistoryRating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.HistoryRating
// @Router /HistoryRatings [post]
func CreateHistoryRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input HistoryRatingInput
	var rating models.Rating
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RatingID not found!"})
		return
	}

	// Create HistoryRating
	HistoryRating := models.HistoryRating{RatingID: input.RatingID, GameID: input.GameID, UserID: input.UserID}
	db.Create(&HistoryRating)

	c.JSON(http.StatusOK, gin.H{"data": HistoryRating})
}

// GetHistoryRatingById godoc
// @Summary Get HistoryRating.
// @Description Get a HistoryRating by id.
// @Tags HistoryRating
// @Produce json
// @Param id path string true "HistoryRating id"
// @Success 200 {object} models.HistoryRating
// @Router /HistoryRatings/{id} [get]
func GetHistoryRatingById(c *gin.Context) { // Get model if exist
	var HistoryRating models.HistoryRating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&HistoryRating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": HistoryRating})
}

//GetHistoryRatingByGame godoc
// @Summary  Get HistoryRating By Game.
// @Description Get HistoryRating by Game.
// @Tags HistoryRating
// @Produce json
// @Param id path string true "Game id"
// @Success 200 {object} []models.Game
// @Router /HistoryRatings/{id}/Games [get]
func GetHistoryRatingByGame(c *gin.Context) { // Get model if exist
	var Games []models.HistoryRating

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("game_id = ?", c.Param("id")).Find(&Games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Games})
}

// UpdateHistoryRating godoc
// @Summary Update HistoryRating.
// @Description Update HistoryRating by id.
// @Tags HistoryRating
// @Produce json
// @Param id path string true "HistoryRating id"
// @Param Body body HistoryRatingInput true "the body to update an HistoryRating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.HistoryRating
// @Router /HistoryRatings/{id} [patch]
func UpdateHistoryRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var HistoryRating models.HistoryRating
	var rating models.HistoryRating
	if err := db.Where("id = ?", c.Param("id")).First(&HistoryRating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input HistoryRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RatingID not found!"})
		return
	}

	var updatedInput models.HistoryRating
	updatedInput.GameID = input.GameID
	updatedInput.RatingID = input.RatingID
	updatedInput.UserID = input.UserID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&HistoryRating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": HistoryRating})
}

// DeleteHistoryRating godoc
// @Summary Delete one HistoryRating.
// @Description Delete a HistoryRating by id.
// @Tags HistoryRating
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "HistoryRating id"
// @Success 200 {object} map[string]boolean
// @Router /HistoryRating/{id} [delete]
func DeleteHistoryRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var HistoryRating models.HistoryRating
	if err := db.Where("id = ?", c.Param("id")).First(&HistoryRating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&HistoryRating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
