package signup

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

type mockUsernameRepository struct {
	insertErr error

	findV   *FindUsernameResponse
	findErr error
}

func (mr *mockUsernameRepository) Insert(username string) error {
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

func TestService_InvalidEmail(t *testing.T) {
	var (
		userRepository     = &mockUserRepository{}
		usernameRepository = &mockUsernameRepository{}
		payload            = CreateUserRequest{
			Username:        "validUsername",
			Email:           "invalidexample.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrInvalidEmail
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_InvalidPassword(t *testing.T) {
	var (
		userRepository     = &mockUserRepository{}
		usernameRepository = &mockUsernameRepository{}
		payload            = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@email.com",
			Password:        "invld",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrInvalidPassword
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_InvalidConfirmPassword(t *testing.T) {
	var (
		userRepository     = &mockUserRepository{}
		usernameRepository = &mockUsernameRepository{}
		payload            = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "invld",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrPasswordNotMatch
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_DuplicateUsername(t *testing.T) {
	var (
		userRepository = &mockUserRepository{
			createV:   nil,
			createErr: errors.New(""),
		}
		usernameRepository = &mockUsernameRepository{
			insertErr: nil,
			findV: &FindUsernameResponse{
				UID:      "mockUID",
				Username: "validUsername",
			},
			findErr: nil,
		}
		payload = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrDuplicateUser
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_UserRepositoryReturnsError(t *testing.T) {
	var (
		userRepository = &mockUserRepository{
			createV:   nil,
			createErr: ErrCreateUser,
		}
		usernameRepository = &mockUsernameRepository{
			insertErr: nil,
			findV:     nil,
			findErr:   sql.ErrNoRows,
		}
		payload = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrCreateUser
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_UsernameRepositoryReturnsNoError(t *testing.T) {
	var (
		userRepository = &mockUserRepository{
			createV: &CreateUserResponse{
				UID:           "mockuid",
				Email:         "fulan@example.com",
				EmailVerified: false,
				PhoneNumber:   "123",
				DisplayName:   "mcok",
				PhotoURL:      "url",
				Disabled:      false,
			},
			createErr: nil,
		}
		usernameRepository = &mockUsernameRepository{
			insertErr: &pq.Error{Code: "23505"},
		}
		payload = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrDuplicateUser
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_NoError(t *testing.T) {
	var (
		userRepository = &mockUserRepository{
			createV: &CreateUserResponse{
				UID:           "mockuid",
				Email:         "fulan@example.com",
				EmailVerified: false,
				PhoneNumber:   "123",
				DisplayName:   "mcok",
				PhotoURL:      "url",
				Disabled:      false,
			},
			createErr: nil,
		}
		usernameRepository = &mockUsernameRepository{
			insertErr: nil,
			findV:     nil,
			findErr:   sql.ErrNoRows,
		}
		payload = CreateUserRequest{
			Username:        "validUsername",
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal = &CreateUserResponse{
			UID:           "mockuid",
			Email:         "fulan@example.com",
			EmailVerified: false,
			PhoneNumber:   "123",
			DisplayName:   "mcok",
			PhotoURL:      "url",
			Disabled:      false,
		}
		expErr error = nil
	)

	s := NewService(context.Background(), userRepository, usernameRepository)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}
