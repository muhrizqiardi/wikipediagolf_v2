package middleware

import "time"

type FindByUIDResponse struct {
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
