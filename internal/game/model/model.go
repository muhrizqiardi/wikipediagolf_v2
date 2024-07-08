package model

import (
	"time"

	"github.com/google/uuid"
)

type Summary struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type Game struct {
	ID         uuid.UUID `db:"id"`
	RoomID     uuid.UUID `db:"room_id"`
	Index      int       `db:"index"`
	IsFinished bool      `db:"is_finished"`
	Language   string    `db:"language"`
	FromTitle  string    `db:"from_title"`
	ToTitle    string    `db:"to_title"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
