package signin

import "time"

type mockRepository struct {
	verifyIDTokenV   *VerifyIDTokenResponse
	verifyIDTokenErr error
	sessionCookieV   *SignInResponse
	sessionCookieErr error
}

func (mr *mockRepository) VerifyIDToken(idTokens string) (*VerifyIDTokenResponse, error) {
	return mr.verifyIDTokenV, mr.verifyIDTokenErr
}

func (mr *mockRepository) SessionCookie(uid string, expiresIn time.Duration) (*SignInResponse, error) {
	return mr.sessionCookieV, mr.sessionCookieErr
}
