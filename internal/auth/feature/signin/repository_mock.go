package signin

import "time"

type mockRepository struct {
	sessionCookieV   *SignInResponse
	sessionCookieErr error
}

func (mr *mockRepository) SessionCookie(uid string, expiresIn time.Duration) (*SignInResponse, error) {
	return mr.sessionCookieV, mr.sessionCookieErr
}
