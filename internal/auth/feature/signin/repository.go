package signin

import (
	"context"
	"log/slog"
	"time"

	firebase "firebase.google.com/go/v4"
)

type Repository interface {
	SessionCookie(uid string, expiresIn time.Duration) (*SignInResponse, error)
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

func (r *repository) SessionCookie(idToken string, expiresIn time.Duration) (*SignInResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	cookie, err := client.SessionCookie(r.context, idToken, expiresIn)
	if err != nil {
		slog.Error("failed to create session cookie", "err", err)
		return nil, err
	}

	return &SignInResponse{
		SessionCookie: cookie,
	}, nil
}
