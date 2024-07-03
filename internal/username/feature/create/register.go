package create

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"
)

func BuildCreate(ctx context.Context, db *sql.DB, tmpl *template.Template, serveMux *http.ServeMux) {
	r := newRepository(ctx, db)
	s := newService(ctx, r)
	addTemplate(tmpl)
	deps := endpointDeps{
		Template: tmpl,
		Service:  s,
	}
	addEndpoint(serveMux, deps)
}
