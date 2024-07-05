package check

type Service interface {
	Check(uid string) error
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Check(uid string) error {
	if _, err := s.repository.GetRoomBelongToMember(uid); err != nil {
		return err
	}

	return nil
}
