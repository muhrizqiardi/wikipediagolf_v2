package create

import "context"

type Service interface {
	Create(payload CreateUsernameRequest) error
}

type service struct {
	context    context.Context
	repository Repository
}

func NewService(ctx context.Context, r Repository) *service {
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

	return s.repository.Insert(payload.UID, payload.Username)
}
