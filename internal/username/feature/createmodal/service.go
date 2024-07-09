package createmodal

import (
	"database/sql"
	"errors"
)

type Service interface {
	FindByUID(uid string) (*FindByUIDResponse, error)
}

type service struct {
	repository Repository
}

func newService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindByUID(uid string) (*FindByUIDResponse, error) {
	u, err := s.repository.FindByUID(uid)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrUsernameNotFound
		}

		return nil, err
	}

	return (*FindByUIDResponse)(u), nil
}
