package createmodal

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"
)

func Register(ctx context.Context, db *sql.DB, tmpl *template.Template, serveMux *http.ServeMux) {
	r := newRepository(ctx, db)
	s := newService(r)
	addTemplate(tmpl)
	deps := endpointDeps{
		Service:  s,
		Template: tmpl,
	}
	addEndpoint(serveMux, deps)
}
