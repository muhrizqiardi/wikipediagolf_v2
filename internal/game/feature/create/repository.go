package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type Repository interface {
	GetRoomBelongToMember(userUID string) (*model.Room, error)
	GetRandomSummary(language string) (*model.Summary, error)
	GetLatestGame(roomID uuid.UUID) (*model.Game, error)
	CreateGame(roomID uuid.UUID, index int, language, fromTitle, toTitle string) (*model.Game, error)
}
