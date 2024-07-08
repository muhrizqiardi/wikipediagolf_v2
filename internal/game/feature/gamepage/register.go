package gamepage

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/repository"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux, ctx context.Context, httpClient *http.Client, db *sql.DB) {
	addTemplate(tmpl)
	r := repository.NewRepository(ctx, httpClient, db)
	s := newService(r)
	c := authcontext.NewAuthContext()
	deps := EndpointDeps{
		Template:    tmpl,
		AuthContext: c,
		Service:     s,
	}
	addEndpoint(serveMux, deps)
}
