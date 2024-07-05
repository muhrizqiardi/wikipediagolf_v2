package check

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"

	firebase "firebase.google.com/go/v4"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/repository"
)

func Register(ctx context.Context, db *sql.DB, firebaseApp *firebase.App, tmpl *template.Template, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, db, firebaseApp)
	addTemplate(tmpl)
	s := newService(r)
	deps := endpointDeps{
		Template:    tmpl,
		Service:     s,
		AuthContext: authcontext.NewAuthContext(),
	}
	addEndpoint(serveMux, deps)
}
