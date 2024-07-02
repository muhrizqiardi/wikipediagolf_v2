package create

import (
	"errors"
	"regexp"
)

var (
	UsernamePattern = regexp.MustCompile("^[A-Za-z0-9_-]+$")

	ErrDuplicateUsername = errors.New("provided username already exist")
	ErrInvalidUsername   = errors.New("invalid username")
	ErrCreateUsername    = errors.New("failed to create username")
)
