package signin

import "time"

type mockService struct {
	signInV   *SignInResponse
	signInErr error
}

func (ms *mockService) SignIn(idToken string, expiresIn time.Duration) (*SignInResponse, error) {
	return ms.signInV, ms.signInErr
}
