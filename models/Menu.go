package models

import "time"

type Menu struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt   time.Time `json:"updated_at" sql:"default:now()"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	DrinkShopID int32     `json:"drink_shop"`
	UserID      int32     `json:"user" pg:"fk:user_id"`
	Orders      []*Order  `json:"orders"`
}

// GetEndTime return endtime unix
func (menu Menu) GetEndTime() int64 {
	return menu.EndTime.Unix()
}
