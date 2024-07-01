package signup

import (
	"context"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

type mockUsernameRepository struct {
	insertErr error

	findV   *FindUsernameResponse
	findErr error
}

func (mr *mockUsernameRepository) Insert(uid, username string) error {
	return mr.insertErr
}

func (mr *mockUsernameRepository) Find(username string) (*FindUsernameResponse, error) {
	return mr.findV, mr.findErr
}

type mockUserRepository struct {
	createV   *CreateUserResponse
	createErr error
}

func (mr *mockUserRepository) Create(email, password string) (*CreateUserResponse, error) {
	return mr.createV, mr.createErr
}

func TestService_InvalidUsername(t *testing.T) {
	var (
		userRepository     = &mockUserRepository{}
		usernameRepository = &mockUsernameRepository{}
		payload            = CreateUserRequest{
			Username:        "invalid/username",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrInvalidUsername
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}
