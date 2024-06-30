package signup

import (
	"errors"
	"regexp"
)

var (
	UsernamePattern = regexp.MustCompile("^[A-Za-z0-9_-]+$")

	ErrCreateUser       = errors.New("failed to create user")
	ErrDuplicateUser    = errors.New("user already exist")
	ErrInvalidUsername  = errors.New("invalid username")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrPasswordNotMatch = errors.New("confirm password does not match password")
	ErrInvalidPassword  = errors.New("password should have at least 8 characters")
)
