package models

import (
	"time"
)

type (
	// review
	Review struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Text      string    `json:"text"`
		GameID    int       `json:"game_id"`
		UserID    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-"`
		Game      Game      `json:"-"`
	}
)
