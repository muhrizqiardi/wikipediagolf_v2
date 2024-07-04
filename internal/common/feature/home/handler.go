package home

import (
	"html/template"
	"log/slog"
	"net/http"

	authctx "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(
	tmpl *template.Template,
	c authctx.AuthContext,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, ok := c.GetFromRequest(r)
		if !ok {
			data := TemplateData{
				IsAuthenticated: false,
			}
			if err := ExecuteTemplate(tmpl, w, data); err != nil {
				slog.Error("failed to execute template", "err", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		data := TemplateData{
			IsAuthenticated: ok,
		}

		if ok {
			data.UID = v.UID
		}

		if err := ExecuteTemplate(tmpl, w, data); err != nil {
			slog.Error("failed to execute template", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	})
}
