package create

import (
	"log/slog"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(s Service, c authcontext.AuthContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, ok := c.GetFromRequest(r)
		if !ok {
			slog.Error("no UID found in request context")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if _, err := s.Create(v.UID); err != nil {
			slog.Error("failed to create new room", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Push-URL", "/rooms")
	})
}
