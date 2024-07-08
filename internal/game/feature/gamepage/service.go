package gamepage

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"

type Service interface {
	CurrentGame(userUID string) (*model.Game, error)
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}
func (s *service) CurrentGame(userUID string) (*model.Game, error) {
	room, err := s.repository.GetRoomBelongToMember(userUID)
	if err != nil {
		return nil, err
	}

	return s.repository.GetLatestGame(room.ID)
}
