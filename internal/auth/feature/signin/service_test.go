package signin

import (
	"errors"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestSignIn_ErrSignIn(t *testing.T) {
	var (
		mr = &mockRepository{
			sessionCookieV:   nil,
			sessionCookieErr: errors.New(""),
		}
		s           = newService(mr)
		mockIDToken = "mockIDToken"
	)

	response, err := s.SignIn(mockIDToken, SessionCookieExpiresDuration)
	testutil.AssertError(t, err)
	testutil.AssertNotNil(t, response)
}
