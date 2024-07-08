package create

import (
	"context"
	"database/sql"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/repository"
)

func Register(ctx context.Context, httpClient *http.Client, db *sql.DB, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, httpClient, db)
	s := newService(r)
	c := authcontext.NewAuthContext()
	deps := endpointDeps{
		Service:     s,
		AuthContext: c,
	}
	addEndpoint(serveMux, deps)
}
