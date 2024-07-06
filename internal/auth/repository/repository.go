package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
)

type Repository struct {
	context     context.Context
	firebaseApp *firebase.App
}

func NewRepository(ctx context.Context, firebaseApp *firebase.App) *Repository {
	return &Repository{
		context:     ctx,
		firebaseApp: firebaseApp,
	}
}

type VerifySessionCookieResponse struct {
	AuthTime int64
	Issuer   string
	Audience string
	Expires  int64
	IssuedAt int64
	Subject  string
	UID      string
	Claims   map[string]interface{}
}

func (r *Repository) VerifySessionCookie(sessionCookie string) (*VerifySessionCookieResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	decoded, err := client.VerifySessionCookie(r.context, sessionCookie)
	if err != nil {
		return nil, err
	}

	response := &VerifySessionCookieResponse{
		AuthTime: decoded.AuthTime,
		Issuer:   decoded.Issuer,
		Audience: decoded.Audience,
		Expires:  decoded.Expires,
		IssuedAt: decoded.IssuedAt,
		Subject:  decoded.Subject,
		UID:      decoded.UID,
		Claims:   decoded.Claims,
	}

	return response, nil
}

type GetUserResponse struct {
	ProviderID             string `json:"providerId,omitempty"`
	UID                    string `json:"rawId,omitempty"`
	DisplayName            string `json:"displayName,omitempty"`
	Email                  string `json:"email,omitempty"`
	PhoneNumber            string `json:"phoneNumber,omitempty"`
	PhotoURL               string `json:"photoUrl,omitempty"`
	CustomClaims           map[string]interface{}
	Disabled               bool
	EmailVerified          bool
	TokensValidAfterMillis int64 // milliseconds since epoch.
	TenantID               string
}

func (r *Repository) GetUser(uid string) (*GetUserResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	u, err := client.GetUser(r.context, uid)
	if err != nil {
		return nil, err
	}

	response := &GetUserResponse{
		ProviderID:             u.ProviderID,
		UID:                    u.UID,
		DisplayName:            u.DisplayName,
		Email:                  u.Email,
		PhoneNumber:            u.PhoneNumber,
		PhotoURL:               u.PhotoURL,
		CustomClaims:           u.CustomClaims,
		Disabled:               u.Disabled,
		EmailVerified:          u.EmailVerified,
		TokensValidAfterMillis: u.TokensValidAfterMillis,
		TenantID:               u.TenantID,
	}

	return response, nil
}
