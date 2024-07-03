package signup

import (
	"errors"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
)

func handleError(tmpl *template.Template, err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var msg string = "Failed to create user"
		switch {
		case errors.Is(err, ErrInvalidUsername):
			msg = "Invalid username format"
		case errors.Is(err, ErrInvalidEmail):
			msg = "Invalid email format"
		case errors.Is(err, ErrInvalidPassword):
			msg = "Invalid password format"
		case errors.Is(err, ErrPasswordNotMatch):
			msg = "Confirm Password should be the same as Password"
		}

		if err := ExecuteTemplate(tmpl, w, TemplateData{"error", msg}); err != nil {
			slog.Error("failed to execute template", "err", err)
		}
	})
}

func handler(s Service, tmpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			slog.Error("failed to parse form", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var payload CreateUserRequest
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			slog.Error("failed to decode body", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if _, err := s.SignUp(payload); err != nil {
			slog.Error("failed to sign up", "err", err)
			handleError(tmpl, err).ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
