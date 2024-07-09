package signup

import (
	"context"

	"firebase.google.com/go/v4/errorutils"
)

type Service interface {
	SignUp(payload CreateUserRequest) (*CreateUserResponse, error)
}

type service struct {
	context        context.Context
	userRepository Repository
}

func newService(ctx context.Context, ur Repository) *service {
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
		if errorutils.IsAlreadyExists(err) {
			return nil, ErrDuplicateEmail
		}

		return nil, err
	}

	return (*CreateUserResponse)(u), nil
}
