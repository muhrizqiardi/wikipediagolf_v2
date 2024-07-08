package gamepage

import (
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

type EndpointDeps struct {
	Template    *template.Template
	AuthContext authcontext.AuthContext
	Service     Service
}

func addEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/game", Handler(deps.Template, deps.AuthContext, deps.Service))
}
