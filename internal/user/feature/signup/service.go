package signup

import (
	"context"
	"fmt"
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
			fmt.Printf("fields: %v\n", fields)
			if _, ok := fields["Username"]; ok {
				fmt.Printf("ok: %v\n", ok)
				return nil, ErrInvalidUsername
			} else if _, ok := fields["Password"]; ok {
				fmt.Printf("ok: %v\n", ok)
				return nil, ErrInvalidPassword
			} else if _, ok := fields["ConfirmPassword"]; ok {
				fmt.Printf("ok: %v\n", ok)
				return nil, ErrPasswordNotMatch
			}
		}
	}

	if _, err := s.usernameRepository.Find(payload.Username); err == nil {
		return nil, err
	}

	return nil, nil
}
