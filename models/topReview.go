package models

import (
	"time"
)

type (
	// topReview
	TopReview struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"name"`
		ReviewID  uint      `json:"review_ID"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Review    Review    `json:"-"`
	}
)
