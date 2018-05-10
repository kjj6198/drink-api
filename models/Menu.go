package models

import "time"

type Menu struct {
	Name      string    `json:"name" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"created_at"`
}
