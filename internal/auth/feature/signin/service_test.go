package signin

import (
	"errors"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestSignIn_ErrSignIn(t *testing.T) {
	var (
		mr = &mockRepository{
			verifyIDTokenV: &VerifyIDTokenResponse{
				UID: "mockUID",
			},
			verifyIDTokenErr: nil,
			sessionCookieV:   nil,
			sessionCookieErr: errors.New(""),
		}
		s           = NewService(mr)
		mockIDToken = "mockIDToken"
	)

	response, err := s.SignIn(mockIDToken, SessionCookieExpiresDuration)
	testutil.AssertError(t, err)
	testutil.AssertNotNil(t, response)
}
