package waitingpage

import (
	"html/template"
	"log/slog"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(
	tmpl *template.Template,
	c authcontext.AuthContext,
	s Service,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, ok := c.GetFromRequest(r)
		if !ok {
			slog.Error("no UID found on request context")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		response, err := s.GetRoom(v.UID)
		if err != nil {
			slog.Error("failed to get room", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var currentUserIsOwner bool
		for _, member := range response.Members {
			if member.UserUID == v.UID {
				slog.Debug("currentUserIsOwner", "currentUserIsOwner", currentUserIsOwner)
				currentUserIsOwner = member.IsOwner
			}
		}

		if err := ExecuteTemplate(tmpl, w, TemplateData{
			CurrentUserUID:     v.UID,
			CurrentUserIsOwner: currentUserIsOwner,
			Room:               response.Room,
			Members:            response.Members,
			MembersTotal:       len(response.Members),
		}); err != nil {
			slog.Error("failed to execute template", "err", err)
			return
		}
	})
}
