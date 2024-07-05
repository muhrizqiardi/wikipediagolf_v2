package check

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"

type Repository interface {
	GetRoomBelongToMember(userUID string) (*model.Room, error)
}
