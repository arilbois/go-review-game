package models

import (
	"time"
)

type (
	// rating
	Rating struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Value     string    `json:"value"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
