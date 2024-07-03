package signin

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
)

func Register(ctx context.Context, firebaseApp *firebase.App, serveMux *http.ServeMux) {
	r := newRepository(ctx, firebaseApp)
	s := newService(r)
	deps := endpointDeps{
		service: s,
	}
	addEndpoint(serveMux, deps)
}
