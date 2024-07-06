package nicknamedialog

import (
	"html/template"
	"log/slog"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func handler(
	tmpl *template.Template,
	_ authcontext.AuthContext,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formType := r.PostFormValue("type")
		if err := executeTemplate(tmpl, w, templateData{
			Type: formType,
		}); err != nil {
			slog.Error("failed to execute template", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
