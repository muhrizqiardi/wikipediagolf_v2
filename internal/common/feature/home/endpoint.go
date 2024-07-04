package home

import (
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

type endpointDeps struct {
	Template    *template.Template
	AuthContext authcontext.AuthContext
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("GET /{$}", handler(deps.Template, deps.AuthContext))
}
