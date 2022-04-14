package controllers

import (
	"net/http"
	"time"

	"final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameInput struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	CategoryID uint   `json:"category_id"`
}

// GetAllGames godoc
// @Summary Get all Games.
// @Description Get a list of Games.
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /Games [get]
func GetAllGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Games []models.Game
	db.Find(&Games)

	c.JSON(http.StatusOK, gin.H{"data": Games})
}

// CreateGame godoc
// @Summary Create New Game.
// @Description Creating a new Game.
// @Tags Game
// @Param Body body GameInput true "the body to create a new Game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /Games [post]
func CreateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input GameInput
	var rating models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.CategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryID not found!"})
		return
	}

	// Create Game
	Game := models.Game{Name: input.Name, Year: input.Year, CategoryID: input.CategoryID}
	db.Create(&Game)

	c.JSON(http.StatusOK, gin.H{"data": Game})
}

// GetGameById godoc
// @Summary Get Game.
// @Description Get a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "Game id"
// @Success 200 {object} models.Game
// @Router /Games/{id} [get]
func GetGameById(c *gin.Context) { // Get model if exist
	var Game models.Game

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Game})
}

// UpdateGame godoc
// @Summary Update Game.
// @Description Update Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "Game id"
// @Param Body body GameInput true "the body to update an Game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Game
// @Router /Games/{id} [patch]
func UpdateGame(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Game models.Game
	var rating models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&Game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input GameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.CategoryID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryID not found!"})
		return
	}

	var updatedInput models.Game
	updatedInput.Name = input.Name
	updatedInput.Year = input.Year
	updatedInput.CategoryID = input.CategoryID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&Game).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Game})
}

//GetGameByCategory godoc
// @Summary  Get Game By Categori.
// @Description Get Game by Categori.
// @Tags Game
// @Produce json
// @Param id path string true "Categori id"
// @Success 200 {object} []models.Category
// @Router /Games/{id}/Categories [get]
func GetGameByCategory(c *gin.Context) { // Get model if exist
	var Category []models.Game

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("category_id = ?", c.Param("id")).Find(&Category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Category})
}

// DeleteGame godoc
// @Summary Delete one Game.
// @Description Delete a Game by id.
// @Tags Game
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game id"
// @Success 200 {object} map[string]boolean
// @Router /Games/{id} [delete]
func DeleteGame(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&Game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
