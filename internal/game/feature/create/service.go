package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
)

type Service interface {
	Create(language, userUID string, roomID uuid.UUID) (*model.Game, error)
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(language, userUID string, roomID uuid.UUID) (*model.Game, error) {
	fromSummary, err := s.repository.GetRandomSummary(language)
	if err != nil {
		return nil, err
	}
	toSummary, err := s.repository.GetRandomSummary(language)
	if err != nil {
		return nil, err
	}

	// TODO: check if room belongs to user UID

	index := 0
	latestGame, err := s.repository.GetLatestGame(roomID)
	if err == nil {
		index = latestGame.Index + 1
	}

	return s.repository.CreateGame(roomID, index, fromSummary.Title, toSummary.Title)
}
