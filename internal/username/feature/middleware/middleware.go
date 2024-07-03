package middleware

import (
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func Middleware(s Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/usernames/create" {
				next.ServeHTTP(w, r)
				return
			}
			val, ok := authcontext.GetFromRequest(r)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			if _, err := s.FindByUID(val.UID); err != nil {
				http.Redirect(w, r, "/usernames/create", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
