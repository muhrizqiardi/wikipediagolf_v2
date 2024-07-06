package middleware

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"

type Repository interface {
	VerifySessionCookie(sessionCookie string) (*repository.VerifySessionCookieResponse, error)
	GetUser(uid string) (*repository.GetUserResponse, error)
}
