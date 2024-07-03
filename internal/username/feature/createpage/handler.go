package createpage

import (
	"html/template"
	"log/slog"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func Handler(tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, _ := authcontext.GetFromRequest(r)
		data := TemplateData{
			UID: val.UID,
		}
		if err := ExecuteTemplate(tmpl, w, data); err != nil {
			slog.Error("failed to execute template", "err", err)
			return
		}
	})
}
