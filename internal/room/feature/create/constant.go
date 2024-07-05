package create

import "errors"

var (
	ErrDuplicateRoomCode = errors.New("room with provided code exists")
)
