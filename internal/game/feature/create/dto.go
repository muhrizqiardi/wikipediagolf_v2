package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type CreateGameRequest struct {
	RoomID uuid.UUID `schema:"roomId"`
}

type CreateGameResponse struct {
	model.Game
}
