package models

import (
	"drink-api/model"
	"time"
)

// User is user
type User struct {
	model.Model
	Email             string    `json:"email,omitempty" binding:"required"`
	EncryptedPassword string    `json:"-"`
	SinginCount       int32     `json:"-"`
	CurrentSignInAt   time.Time `json:"-"`
	Picture           string    `json:"picture,omitempty"`
	UserName          string    `json:"user_name,omitempty"`
	IsAdmin           bool      `json:"is_admin"`
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
