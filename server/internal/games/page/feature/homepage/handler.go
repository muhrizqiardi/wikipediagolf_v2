package homepage

import (
	"html/template"
	"log/slog"
	"net/http"
)

func Endpoint(
	tmpl *template.Template,
) func(serveMux *http.ServeMux) {
	return func(serveMux *http.ServeMux) {
		serveMux.Handle("/{$}", Handler(tmpl))
	}
}

func Handler(
	tmpl *template.Template,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			slog.Error("failed to execute template")
		}
	})
}
