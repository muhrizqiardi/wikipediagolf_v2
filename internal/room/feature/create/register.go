package create

import (
	"context"
	"database/sql"
	"net/http"

	firebase "firebase.google.com/go/v4"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/repository"
)

func Register(ctx context.Context, db *sql.DB, firebaseApp *firebase.App, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, db)
	s := NewService(NewCodeGenerator(), r)
	deps := endpointDeps{
		Service:     s,
		AuthContext: authcontext.NewAuthContext(),
	}
	addEndpoint(serveMux, deps)
}
