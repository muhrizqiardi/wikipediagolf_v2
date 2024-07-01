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
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
