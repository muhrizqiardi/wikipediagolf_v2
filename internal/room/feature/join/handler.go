package join

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(s Service, c authcontext.AuthContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var payload JoinRequest
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			slog.Error("failed to decode", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		v, ok := c.GetFromRequest(r)
		if !ok {
			slog.Error("no UID found in request context")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if err := s.Join(v.UID, payload.RoomCode); err != nil {
			slog.Error("failed to join room", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("HX-Redirect", "/rooms")
	})
}
