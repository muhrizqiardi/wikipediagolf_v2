package authmiddleware

import (
	"log/slog"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/authctx"
)

func AuthMiddleware(firebaseApp *firebase.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			client, err := firebaseApp.Auth(r.Context())
			if err != nil {
				slog.Warn("failed to instantiate Firebase auth client", "err", err)
				next.ServeHTTP(w, r)
				return
			}

			decoded, err := client.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			authctx.SetRequest(r, authctx.Val{
				UID: decoded.UID,
			})

			next.ServeHTTP(w, r)
		})
	}
}