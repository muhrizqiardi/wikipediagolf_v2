package createmodal

import (
	"html/template"
	"log/slog"
	"net/http"

	authctx "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(s Service, tmpl *template.Template, c authctx.AuthContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, ok := c.GetFromRequest(r)
		if !ok {
			w.WriteHeader(http.StatusOK)
			return
		}

		if _, err := s.FindByUID(v.UID); err != nil {
			if err := ExecuteTemplate(tmpl, w, TemplateData{
				UID: v.UID,
			}); err != nil {
				slog.Error("failed to execute template", "err", err)
				return
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
