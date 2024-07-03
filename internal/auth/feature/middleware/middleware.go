package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func AuthMiddleware(firebaseApp *firebase.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerHeader := r.Header.Get("Bearer")
			bearerHeaderSplit := strings.Split(bearerHeader, " ")
			if len(bearerHeaderSplit) < 2 {
				next.ServeHTTP(w, r)
				return
			}
			idToken := bearerHeaderSplit[1]

			client, err := firebaseApp.Auth(r.Context())
			if err != nil {
				slog.Warn("failed to instantiate Firebase auth client", "err", err)
				next.ServeHTTP(w, r)
				return
			}

			decoded, err := client.VerifyIDToken(r.Context(), idToken)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			authcontext.SetRequest(r, authcontext.Val{
				UID: decoded.UID,
			})

			next.ServeHTTP(w, r)
		})
	}
}
