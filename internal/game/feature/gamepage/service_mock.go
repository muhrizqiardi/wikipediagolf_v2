package gamepage

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"

type mockService struct {
	v   *model.Game
	err error
}

func (ms *mockService) CurrentGame(userUID string) (*model.Game, error) {
	return ms.v, ms.err
}
