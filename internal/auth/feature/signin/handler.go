package signin

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
)

func Handler(
	s Service,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var sessionToken SignInRequest
		if err := schema.NewDecoder().Decode(&sessionToken, r.PostForm); err != nil {
			slog.Error("failed to decode body", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cookie, err := s.SignIn(sessionToken.IDToken, SessionCookieExpiresDuration)
		if err != nil {
			slog.Error("failed to sign in", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    cookie.SessionCookie,
			MaxAge:   int(SessionCookieExpiresDuration.Seconds()),
			HttpOnly: true,
		})

		w.WriteHeader(http.StatusOK)
	})
}
