package middleware

import (
	"log/slog"
	"net/http"

	firebase "firebase.google.com/go/v4"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func AuthMiddleware(
	firebaseApp *firebase.App,
	c authcontext.AuthContext,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client, err := firebaseApp.Auth(r.Context())
			if err != nil {
				slog.Warn("failed to instantiate Firebase auth client", "err", err)
				next.ServeHTTP(w, r)
				return
			}
			cookie, err := r.Cookie("session")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			decoded, err := client.VerifySessionCookie(r.Context(), cookie.Value)
			if err != nil {
				slog.Debug("verify session cookie failed", "err", err)
				next.ServeHTTP(w, r)
				return
			}

			c.SetRequest(r, authcontext.Val{
				UID: decoded.UID,
			})

			next.ServeHTTP(w, r)
		})
	}
}
