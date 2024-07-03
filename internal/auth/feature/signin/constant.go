package signin

import (
	"errors"
	"time"
)

var (
	ErrSignIn = errors.New("failed to sign in")
)

const (
	SessionCookieExpiresDuration = time.Hour * 24 * 5
)
