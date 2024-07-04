package waitingpage

import (
	"time"

	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
)

type GetRoomResponseMember struct {
	ID          uuid.UUID
	IsOwner     bool
	RoomID      uuid.UUID
	UserUID     string
	Username    string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetRoomResponse struct {
	Room    model.Room
	Members []GetRoomResponseMember
}
