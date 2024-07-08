package create

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(s Service, c authcontext.AuthContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, _ := c.GetFromRequest(r)
		r.ParseForm()

		var payload CreateGameRequest
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			slog.Error("failed to decode payload", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := s.Create("en", v.UID, payload.RoomID); err != nil {
			slog.Error("failed to create game", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/game")
		w.WriteHeader(http.StatusCreated)
	})
}
