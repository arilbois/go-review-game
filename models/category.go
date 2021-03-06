package models

import (
	"time"
)

type (
	// category
	Category struct {
		ID          uint      `json:"id" gorm:"primary_key"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
