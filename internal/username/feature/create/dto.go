package create

import "time"

type FindUsernameResponse struct {
	UID      string
	Username string
}

type CreateUsernameRequest struct {
	UID      string
	Username string
}

type CreateUsernameResponse struct {
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
