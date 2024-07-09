package create

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/username/repository"
)

func BuildCreate(ctx context.Context, db *sql.DB, tmpl *template.Template, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, db)
	s := newService(ctx, r)
	addTemplate(tmpl)
	deps := endpointDeps{
		Template: tmpl,
		Service:  s,
	}
	addEndpoint(serveMux, deps)
}
