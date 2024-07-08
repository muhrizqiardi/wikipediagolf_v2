package gamepage

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type Repository interface {
	GetLatestGame(roomID uuid.UUID) (*model.Game, error)
	GetRoomBelongToMember(userUID string) (*model.Room, error)
}
