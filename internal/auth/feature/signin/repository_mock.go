package signin

import (
	"time"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"
)

type mockRepository struct {
	sessionCookieV   *repository.SessionCookieResult
	sessionCookieErr error
}

func (mr *mockRepository) SessionCookie(uid string, expiresIn time.Duration) (*repository.SessionCookieResult, error) {
	return mr.sessionCookieV, mr.sessionCookieErr
}
