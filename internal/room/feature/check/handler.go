package check

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
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := s.Check(v.UID); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := executeTemplate(tmpl, w); err != nil {
			slog.Error("failed to execute template", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
