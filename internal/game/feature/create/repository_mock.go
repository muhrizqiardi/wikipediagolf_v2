package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type mockRepository struct {
	getRandomSummaryV   *model.Summary
	getRandomSummaryErr error
	getLatestGameV      *model.Game
	getLatestGameErr    error
	createGameV         *model.Game
	createGameErr       error
}

func (mr *mockRepository) GetRandomSummary(language string) (*model.Summary, error) {
	return mr.getRandomSummaryV, mr.getRandomSummaryErr
}

func (mr *mockRepository) GetLatestGame(roomID uuid.UUID) (*model.Game, error) {
	return mr.getLatestGameV, mr.getLatestGameErr
}

func (mr *mockRepository) CreateGame(roomID uuid.UUID, index int, fromTitle, toTitle string) (*model.Game, error) {
	return mr.createGameV, mr.createGameErr
}
