package createpage

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/authctx"
)

func Handler(tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, ok := authctx.GetFromRequest(r)
		if !ok {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		data := TemplateData{
			UID: val.UID,
		}
		if err := ExecuteTemplate(tmpl, w, data); err != nil {
			slog.Error("failed to execute template", "err", err)
			return
		}
	})
}
