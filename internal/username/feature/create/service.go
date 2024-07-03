package create

import (
	"context"

	"github.com/lib/pq"
)

type Service interface {
	Create(payload CreateUsernameRequest) error
}

type service struct {
	context    context.Context
	repository Repository
}

func newService(ctx context.Context, r Repository) *service {
	return &service{
		context:    ctx,
		repository: r,
	}
}

func (s *service) Create(payload CreateUsernameRequest) error {
	if err := payload.Valid(s.context); err != nil {
		if err, ok := err.(*ValidationErrors); ok {
			fields := err.Fields()
			if _, ok := fields["Username"]; ok {
				return ErrInvalidUsername
			}
		}
	}

	if err := s.repository.Insert(payload.UID, payload.Username); err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch {
			case err.Code == "23505":
				return ErrDuplicateUsername
			}
		}
		return ErrCreateUsername
	}

	return nil
}
