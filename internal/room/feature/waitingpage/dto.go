package waitingpage

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"

type GetRoomResponse struct {
	Room    model.Room
	Members []model.RoomMember
}
