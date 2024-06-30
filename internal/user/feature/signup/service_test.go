package signup

import (
	"context"
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

func TestService(t *testing.T) {
	type testCase struct {
		UserRepository     UserRepository
		UsernameRepository UsernameRepository
		Name               string
		Payload            CreateUserRequest
		ExpVal             *CreateUserResponse
		ExpErr             error
	}
	tests := []testCase{
		{
			Name: "should return error if username is invalid",

			UserRepository:     &mockUserRepository{},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "invalid/username",
				Email:           "valid@example.com",
				Password:        "secure_Password321",
				ConfirmPassword: "secure_Password321",
			},
			ExpVal: nil,
			ExpErr: ErrInvalidUsername,
		},
		{
			Name: "should return error if email is invalid",

			UserRepository:     &mockUserRepository{},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "invalidexample.com",
				Password:        "secure_Password321",
				ConfirmPassword: "secure_Password321",
			},
			ExpVal: nil,
			ExpErr: ErrInvalidEmail,
		},
		{
			Name: "should return error if password is invalid format",

			UserRepository:     &mockUserRepository{},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "invalidexample.com",
				Password:        "invld",
				ConfirmPassword: "",
			},
			ExpVal: nil,
			ExpErr: ErrInvalidPassword,
		},
		{
			Name: "should return error if confirm password does not match password",

			UserRepository:     &mockUserRepository{},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "invalidexample.com",
				Password:        "secure_Password321",
				ConfirmPassword: "invld",
			},
			ExpVal: nil,
			ExpErr: ErrPasswordNotMatch,
		},
		{
			Name: "should return error if user repository returns error",

			UserRepository: &mockUserRepository{
				createV:   nil,
				createErr: errors.New(""),
			},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "valid@example.com",
				Password:        "secure_Password321",
				ConfirmPassword: "secure_Password321",
			},
			ExpVal: nil,
			ExpErr: ErrCreateUser,
		},
		{
			Name: "should return appropriate error if username repository returns unique violation error",

			UserRepository: &mockUserRepository{
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
			},
			UsernameRepository: &mockUsernameRepository{
				insertErr: &pq.Error{Code: "23505"},
			},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "valid@example.com",
				Password:        "secure_Password321",
				ConfirmPassword: "secure_Password321",
			},
			ExpVal: nil,
			ExpErr: ErrDuplicateUser,
		},
		{
			Name: "should create user",

			UserRepository: &mockUserRepository{
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
			},
			UsernameRepository: &mockUsernameRepository{},
			Payload: CreateUserRequest{
				Username:        "validUsername",
				Email:           "valid@example.com",
				Password:        "secure_Password321",
				ConfirmPassword: "secure_Password321"},
			ExpVal: &CreateUserResponse{},
			ExpErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s := NewService(context.Background(), test.UserRepository, test.UsernameRepository)
			got, err := s.SignUp(test.Payload)
			if test.ExpErr != nil {
				testutil.CompareError(t, test.ExpErr, err)
			}
			testutil.AssertEqualCMP(t, test.ExpVal, got)
		})
	}
}
