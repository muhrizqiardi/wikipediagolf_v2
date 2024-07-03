package signup

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/errorutils"
)

type Repository interface {
	Create(email, password string) (*CreateUserResponse, error)
}

type repository struct {
	context     context.Context
	firebaseApp *firebase.App
}

func newRepository(ctx context.Context, firebaseApp *firebase.App) *repository {
	return &repository{
		context:     ctx,
		firebaseApp: firebaseApp,
	}
}

func (r *repository) Create(email, password string) (*CreateUserResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	userToCreate := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		EmailVerified(false).
		Disabled(false)

	newUser, err := client.CreateUser(r.context, userToCreate)
	if err != nil {
		switch {
		case errorutils.IsAlreadyExists(err):
			return nil, ErrDuplicateEmail
		}
		return nil, err
	}

	return &CreateUserResponse{
		UID:           newUser.UID,
		Email:         newUser.Email,
		EmailVerified: newUser.EmailVerified,
		PhoneNumber:   newUser.PhoneNumber,
		DisplayName:   newUser.DisplayName,
		PhotoURL:      newUser.PhotoURL,
		Disabled:      newUser.Disabled,
	}, nil
}
