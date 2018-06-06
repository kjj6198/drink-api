package models

import (
	"time"
)

// DrinkShop is drink shop.
type DrinkShop struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Phone     string    `json:"phone"`
	Comment   string    `json:"comment"`
	Address   string    `json:"address"`
	Rank      int16     `json:"rank"`
	ImageURL  string    `json:"image_url" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
