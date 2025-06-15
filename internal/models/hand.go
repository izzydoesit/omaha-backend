package models

import (
	"time"
)

// Hand model
type Hand struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID string `json:"user_id"`  // if using user auth
	Cards string `json:"cards"`		// Store as comma-separated string or JSON
	CreatedAt time.Time `json:"created_at"`
}