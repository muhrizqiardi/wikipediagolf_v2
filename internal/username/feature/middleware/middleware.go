package middleware

import (
	"net/http"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/authctx"
)

func Middleware(s Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val, ok := authctx.GetFromRequest(r)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			if _, err := s.FindByUID(val.UID); err != nil {
				http.Redirect(w, r, "/usernames/create", http.StatusSeeOther)
			}

			next.ServeHTTP(w, r)
		})
	}
}
