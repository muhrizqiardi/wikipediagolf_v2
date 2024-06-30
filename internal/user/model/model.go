package model

import "time"

type User struct {
	UID           string
	Email         string
	EmailVerified bool
	PhoneNumber   string
	Password      string
	DisplayName   string
	PhotoURL      string
	Disabled      string
}

type Username struct {
	UserID    string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
