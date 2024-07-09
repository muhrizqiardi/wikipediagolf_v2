package signup

import (
	"context"
	"html/template"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"
)

func Register(ctx context.Context, firebaseApp *firebase.App, tmpl *template.Template, serveMux *http.ServeMux) {
	r := repository.NewRepository(ctx, firebaseApp)
	s := newService(ctx, r)
	deps := endpointDeps{
		Service:  s,
		Template: tmpl,
	}
	addTemplate(tmpl)
	addEndpoint(serveMux, deps)
}
