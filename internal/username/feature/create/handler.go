package create

import (
	"errors"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
)

func Handler(tmpl *template.Template, service Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			slog.Error("failed to parse form", "err", err)
			ExecuteTemplate(tmpl, w, TemplateData{"error", "Failed to create username"})
			return
		}
		var payload CreateUsernameRequest
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			slog.Error("failed to decode form values", "err", err)
			ExecuteTemplate(tmpl, w, TemplateData{"error", "Failed to create username"})
			return
		}

		if err := service.Create(payload); err != nil {
			slog.Error("failed to create username", "err", err)
			switch {
			case errors.Is(err, ErrInvalidUsername):
				ExecuteTemplate(tmpl, w, TemplateData{"error", "Invalid username format"})
				return
			case errors.Is(err, ErrDuplicateUsername):
				ExecuteTemplate(tmpl, w, TemplateData{"error", "Username already exists"})
				return
			default:
				ExecuteTemplate(tmpl, w, TemplateData{"error", "Failed to create username"})
				return
			}
		}

		w.Header().Set("HX-Redirect", "/")
		w.Header().Set("HX-Target", "body")
		w.Header().Set("HX-Swap", "outerHTML")
		return
	})
}
