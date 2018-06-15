package models

import (
	"drink-api/model"
)

type Order struct {
	model.Model
	Name    string `json:"name" binding:"required"`
	HasPaid bool   `json:"has_paid"`
	Note    string `json:"note"`
	Menu    *Menu  `json:"menu"`
	User    *User  `json:"user"`
}
