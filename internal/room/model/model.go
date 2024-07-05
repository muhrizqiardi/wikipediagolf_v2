package model

import (
	"time"

	"github.com/google/uuid"
)

type RoomMember struct {
	ID        uuid.UUID `db:"id"`
	IsOwner   bool      `db:"is_owner"`
	RoomID    uuid.UUID `db:"room_id"`
	UserUID   string    `db:"user_uid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Room struct {
	ID        uuid.UUID `db:"id"`
	Code      string    `db:"code"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
