package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
)

type Repository interface {
	InsertRoom(roomCode, status string) (*model.Room, error)
	InsertRoomMember(roomID uuid.UUID, userUID string, isOwner bool) (*model.RoomMember, error)
}
