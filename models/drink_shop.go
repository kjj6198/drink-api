package models

import (
	"drink-api/model"
)

// DrinkShop is drink shop.
type DrinkShop struct {
	model.Model
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
	Comment  string `json:"comment"`
	Address  string `json:"address"`
	Rank     int16  `json:"rank"`
	ImageURL string `json:"image_url" binding:"required"`
}
