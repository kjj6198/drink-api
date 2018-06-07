package models

import (
	"time"
)

type Order struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name" binding:"required"`
	HasPaid   bool      `json:"has_paid"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"default:now()"`
	Menu      *Menu     `json:"menu" pg:"fk:menu_id"`
	User      *User     `json:"user" pg:"fk:user_id"`
}
