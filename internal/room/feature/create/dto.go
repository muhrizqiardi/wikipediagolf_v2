package create

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"

type CreateRoomResponse struct {
	Room  model.Room
	Owner model.RoomMember
}
