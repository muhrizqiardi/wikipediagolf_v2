package createmodal

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func Register(ctx context.Context, db *sql.DB, tmpl *template.Template, serveMux *http.ServeMux, ac authcontext.AuthContext) {
	r := newRepository(ctx, db)
	s := newService(r)
	addTemplate(tmpl)
	deps := endpointDeps{
		Service:     s,
		Template:    tmpl,
		AuthContext: ac,
	}
	addEndpoint(serveMux, deps)
}
