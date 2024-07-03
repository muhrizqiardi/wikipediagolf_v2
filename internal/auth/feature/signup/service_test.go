package signup

import (
	"context"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestService_InvalidEmail(t *testing.T) {
	var (
		mr      = &mockRepository{}
		payload = CreateUserRequest{
			Email:           "invalidexample.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrInvalidEmail
	)

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_InvalidPassword(t *testing.T) {
	var (
		mr      = &mockRepository{}
		payload = CreateUserRequest{
			Email:           "valid@email.com",
			Password:        "invld",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrInvalidPassword
	)

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_InvalidConfirmPassword(t *testing.T) {
	var (
		mr      = &mockRepository{}
		payload = CreateUserRequest{
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "invld",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrPasswordNotMatch
	)

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_DuplicateEmail(t *testing.T) {
	var (
		mr = &mockRepository{
			createV:   nil,
			createErr: ErrDuplicateEmail,
		}
		payload = CreateUserRequest{
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrDuplicateEmail
	)

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_UserRepositoryFindReturnsError(t *testing.T) {
	var (
		mr = &mockRepository{
			createV:   nil,
			createErr: ErrDuplicateEmail,
		}
		payload = CreateUserRequest{
			Email:           "valid@example.com",
			Password:        "secure_Password321",
			ConfirmPassword: "secure_Password321",
		}
		expVal *CreateUserResponse = nil
		expErr                     = ErrDuplicateEmail
	)

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}

func TestService_NoError(t *testing.T) {
	var (
		mr = &mockRepository{
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
		payload = CreateUserRequest{
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

	s := NewService(context.Background(), mr)
	got, err := s.SignUp(payload)
	testutil.CompareError(t, expErr, err)
	testutil.AssertEqualCMP(t, expVal, got)
}
