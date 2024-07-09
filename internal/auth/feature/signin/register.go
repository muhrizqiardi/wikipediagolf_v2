package signin

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"
)

func Register(ctx context.Context, firebaseApp *firebase.App, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, firebaseApp)
	s := newService(r)
	deps := endpointDeps{
		service: s,
	}
	addEndpoint(serveMux, deps)
}
