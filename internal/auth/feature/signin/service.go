package signin

import "time"

type Service interface {
	SignIn(idToken string, expiresIn time.Duration) (*SignInResponse, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) SignIn(idToken string, expiresIn time.Duration) (*SignInResponse, error) {
	decoded, err := s.repository.VerifyIDToken(idToken)
	if err != nil {
		return nil, err
	}
	uid := decoded.UID
	return s.repository.SessionCookie(uid, expiresIn)
}
