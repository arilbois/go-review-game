package models

import (
	"time"
)

type (
	// game
	Game struct {
		ID         uint      `gorm:"primary_key" json:"id"`
		Name       string    `json:"name"`
		Year       int       `json:"year"`
		CategoryID uint      `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Category   Category  `json:"-"`
	}
)
