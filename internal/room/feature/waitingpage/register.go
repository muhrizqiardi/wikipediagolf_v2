package waitingpage

import (
	"context"
	"database/sql"
	"html/template"
	"net/http"

	firebase "firebase.google.com/go/v4"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/repository"
)

func Register(
	ctx context.Context,
	db *sql.DB,
	firebaseApp *firebase.App,
	tmpl *template.Template,
	serveMux *http.ServeMux,
) {
	AddTemplate(tmpl)
	r := repository.NewRepository(ctx, db, firebaseApp)
	deps := EndpointDeps{
		Template:    tmpl,
		AuthContext: authcontext.NewAuthContext(),
		Service:     NewService(r),
	}
	AddEndpoint(serveMux, deps)
}
