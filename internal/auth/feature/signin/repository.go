package signin

import (
	"context"
	"time"

	firebase "firebase.google.com/go/v4"
)

type Repository interface {
	VerifyIDToken(idTokens string) (*VerifyIDTokenResponse, error)
	SessionCookie(uid string, expiresIn time.Duration) (*SignInResponse, error)
}

type repository struct {
	context     context.Context
	firebaseApp *firebase.App
}

func NewRepository(ctx context.Context, firebaseApp *firebase.App) *repository {
	return &repository{
		context:     ctx,
		firebaseApp: firebaseApp,
	}
}

func (r *repository) SessionCookie(idToken string, expiresIn time.Duration) (*SignInResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}
	decoded, err := client.VerifyIDToken(r.context, idToken)
	if err != nil {
		return nil, err
	}
	cookie, err := client.CustomToken(r.context, decoded.UID)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		SessionCookie: cookie,
	}, nil
}
