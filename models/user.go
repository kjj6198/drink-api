package models

import (
	"time"
)

// User is user
type User struct {
	ID                string `json:"id"`
	Email             string `json:"email" binding:"required"`
	EncryptedPassword string
	SinginCount       int32
	CurrentSignInAt   time.Time
	CreatedAt         time.Time `json:"created_at" binding:"required"`
	UpdatedAt         time.Time `json:"updated_at" binding:"required"`
	Picture           string    `json:"picture"`
	UserName          string    `json:"user_name"`
	IsAdmin           bool
	Orders            []*Order `json:"orders"`
	Menus             []*Menu  `json:"menus"`
}

// GoogleAccountInfo is account info from google oauth.
type GoogleAccountInfo struct {
	Email string
	Name  string
	Image string
}

// GetName return username or email
func (user *User) GetName() string {
	if len(user.UserName) > 0 {
		return user.UserName
	}

	return user.Email
}

// FromOauth will receive oauth info and create user
func (user *User) FromOauth(account GoogleAccountInfo) error {
	return nil
}
