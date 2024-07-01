package signup

import (
	"context"
)

type Service interface {
	SignUp(payload CreateUserRequest) (*CreateUserResponse, error)
}

type service struct {
	context            context.Context
	userRepository     UserRepository
	usernameRepository UsernameRepository
}

func NewService(ctx context.Context, ur UserRepository, unr UsernameRepository) *service {
	return &service{
		context:            ctx,
		userRepository:     ur,
		usernameRepository: unr,
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

	if _, err := s.usernameRepository.Find(payload.Username); err == nil {
		return nil, ErrDuplicateUser
	}

	u, err := s.userRepository.Create(payload.Email, payload.Password)
	if err != nil {
		return nil, err
	}

	if err := s.usernameRepository.Insert(u.UID, payload.Username); err != nil {
		return nil, err
	}

	return u, nil
}
