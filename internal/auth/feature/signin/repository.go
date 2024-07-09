package signin

import (
	"time"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"
)

type Repository interface {
	SessionCookie(uid string, expiresIn time.Duration) (*repository.SessionCookieResult, error)
}
