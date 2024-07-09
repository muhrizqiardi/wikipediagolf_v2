package signin

import (
	"time"
)

type Service interface {
	SignIn(idToken string, expiresIn time.Duration) (*SignInResponse, error)
}

type service struct {
	repository Repository
}

func newService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) SignIn(idToken string, expiresIn time.Duration) (*SignInResponse, error) {
	response, err := s.repository.SessionCookie(idToken, expiresIn)
	if err != nil {
		return nil, err
	}

	return (*SignInResponse)(response), nil
}
