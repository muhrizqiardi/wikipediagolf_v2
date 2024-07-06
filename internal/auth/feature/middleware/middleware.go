package middleware

import (
	"log/slog"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func AuthMiddleware(
	s Service,
	c authcontext.AuthContext,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				slog.Debug("cookie not found in request", "err", err)
				next.ServeHTTP(w, r)
				return
			}

			u, err := s.GetUserFromToken(cookie.Value)
			if err != nil {
				slog.Error("user data cannot be retrieved from session cookie value", "err", err)
				next.ServeHTTP(w, r)
				return
			}

			c.SetRequest(r, authcontext.Val{
				UID:    u.UID,
				IsAnon: u.Email == "",
			})

			next.ServeHTTP(w, r)
		})
	}
}
