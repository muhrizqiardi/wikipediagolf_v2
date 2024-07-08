package gamepage

import (
	"html/template"
	"log/slog"
	"net/http"
	"strings"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func Handler(tmpl *template.Template, c authcontext.AuthContext, s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, _ := c.GetFromRequest(r)
		game, err := s.CurrentGame(v.UID)
		if err != nil {
			slog.Error("failed to get current game", "err", err)
			return
		}

		data := templateData{
			FromTitle:        game.FromTitle,
			FromTitleDecoded: strings.ReplaceAll(game.FromTitle, "_", " "),
			ToTitle:          game.ToTitle,
			ToTitleDecoded:   strings.ReplaceAll(game.ToTitle, "_", " "),
		}

		if err := ExecuteTemplate(tmpl, w, data); err != nil {
			slog.Error("failed to execute template", "err", err)
		}
	})
}
