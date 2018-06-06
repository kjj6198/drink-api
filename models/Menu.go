package models

import "time"

type Menu struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name" binding:"required"`
	EndTime   time.Time  `json:"end_time" binding:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DrinkShop *DrinkShop `json:"drink_shop" pg:"fk:drink_shop_id"`
	User      *User      `json:"user" pg:"fk:user_id"`
	Orders    []*Order   `json:"orders"`
}
