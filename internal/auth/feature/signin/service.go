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
	return s.repository.SessionCookie(idToken, expiresIn)
}
