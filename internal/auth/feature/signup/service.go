package signup

import (
	"context"
)

type Service interface {
	SignUp(payload CreateUserRequest) (*CreateUserResponse, error)
}

type service struct {
	context        context.Context
	userRepository Repository
}

func NewService(ctx context.Context, ur Repository) *service {
	return &service{
		context:        ctx,
		userRepository: ur,
	}
}

func (s *service) SignUp(payload CreateUserRequest) (*CreateUserResponse, error) {
	if err := payload.Valid(s.context); err != nil {
		if err, ok := err.(*ValidationErrors); ok {
			fields := err.Fields()
			if _, ok := fields["Username"]; ok {
				return nil, ErrInvalidUsername
			} else if _, ok := fields["Email"]; ok {
				return nil, ErrInvalidEmail
			} else if _, ok := fields["Password"]; ok {
				return nil, ErrInvalidPassword
			} else if _, ok := fields["ConfirmPassword"]; ok {
				return nil, ErrPasswordNotMatch
			}
		}
	}

	u, err := s.userRepository.Create(payload.Email, payload.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}