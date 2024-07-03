package context

import (
	"context"
	"net/http"
)

type Val struct {
	UID string
}

type Key string

const key Key = "user"

func SetRequest(r *http.Request, val Val) {
	req := r.WithContext(
		context.WithValue(r.Context(), key, val),
	)
	*r = *req
}

func GetFromRequest(r *http.Request) (Val, bool) {
	u, ok := (r.Context().Value(key)).(Val)
	return u, ok
}
