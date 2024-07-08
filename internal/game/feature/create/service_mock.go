package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type mockService struct {
	v   *model.Game
	err error
}

func (ms *mockService) Create(language, userUID string, roomID uuid.UUID) (*model.Game, error) {
	return ms.v, ms.err
}
