package middleware

type Service interface {
	GetUserFromToken(token string) (*GetUserResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetUserFromToken(token string) (*GetUserResponse, error) {
	decoded, err := s.repository.VerifySessionCookie(token)
	if err != nil {
		return nil, err
	}

	u, err := s.repository.GetUser(decoded.UID)
	if err != nil {
		return nil, err
	}

	return (*GetUserResponse)(u), err
}
