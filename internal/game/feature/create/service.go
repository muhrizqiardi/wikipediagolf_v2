package create

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"

type Service interface {
	Create(language, userUID string) (*model.Game, error)
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(language, userUID string) (*model.Game, error) {
	fromSummary, err := s.repository.GetRandomSummary(language)
	if err != nil {
		return nil, err
	}
	toSummary, err := s.repository.GetRandomSummary(language)
	if err != nil {
		return nil, err
	}

	room, err := s.repository.GetRoomBelongToMember(userUID)
	if err != nil {
		return nil, err
	}

	index := 0
	latestGame, err := s.repository.GetLatestGame(room.ID)
	if err == nil {
		index = latestGame.Index + 1
	}

	return s.repository.CreateGame(room.ID, index, language, fromSummary.Title, toSummary.Title)
}
