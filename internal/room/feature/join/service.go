package join

type Service interface {
	Join(uid, roomCode string) error
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Join(uid, roomCode string) error {
	room, err := s.repository.GetRoomByCode(roomCode)
	if err != nil {
		return err
	}

	if _, err := s.repository.InsertRoomMember(room.ID, uid, false); err != nil {
		return err
	}

	return nil
}
