package pregame

import (
	"html/template"
	"log/slog"
	"net/http"
)

func handler(tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := ExecuteTemplate(tmpl, w); err != nil {
			slog.Error("failed to execute template", "err", err)
		}
	})
}
