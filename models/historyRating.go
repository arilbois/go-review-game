package models

import (
	"time"
)

type (
	// HistoryRating
	HistoryRating struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		RatingID  uint      `json:"rating_id"`
		GameID    uint      `json:"game_id"`
		UserID    uint      `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-"`
		Rating    Rating    `json:"-"`
		Game      Game      `json:"-"`
	}
)
