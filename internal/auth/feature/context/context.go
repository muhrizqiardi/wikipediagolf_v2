package context

import (
	"context"
	"net/http"
)

type AuthContext interface {
	SetRequest(r *http.Request, val Val)
	GetFromRequest(r *http.Request) (Val, bool)
}

type authContext struct {
}

func NewAuthContext() *authContext {
	return &authContext{}
}

type Val struct {
	UID string
}

type Key string

const key Key = "user"

func (c *authContext) SetRequest(r *http.Request, val Val) {
	req := r.WithContext(
		context.WithValue(r.Context(), key, val),
	)
	*r = *req
}

func (c *authContext) GetFromRequest(r *http.Request) (Val, bool) {
	u, ok := (r.Context().Value(key)).(Val)
	return u, ok
}
