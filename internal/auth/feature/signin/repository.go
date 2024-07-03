package signin

import (
	"context"
	"time"

	firebase "firebase.google.com/go/v4"
)

type Repository interface {
	VerifyIDToken(idToken string) (*VerifyIDTokenResponse, error)
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

func (r *repository) VerifyIDToken(idToken string) (*VerifyIDTokenResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	decoded, err := client.VerifyIDToken(r.context, idToken)
	if err != nil {
		return nil, err
	}

	result := VerifyIDTokenResponse{
		AuthTime: decoded.AuthTime,
		Issuer:   decoded.Issuer,
		Audience: decoded.Audience,
		Expires:  decoded.Expires,
		IssuedAt: decoded.IssuedAt,
		Subject:  decoded.Subject,
		UID:      decoded.UID,
		Firebase: struct {
			SignInProvider string                 "json:\"sign_in_provider\""
			Tenant         string                 "json:\"tenant\""
			Identities     map[string]interface{} "json:\"identities\""
		}{
			SignInProvider: decoded.Firebase.SignInProvider,
			Tenant:         decoded.Firebase.Tenant,
			Identities:     decoded.Firebase.Identities,
		},
		Claims: decoded.Claims,
	}

	return &result, nil
}

func (r *repository) SessionCookie(uid string, expiresIn time.Duration) (*SignInResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	cookie, err := client.CustomToken(r.context, uid)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		SessionCookie: cookie,
	}, nil
}
