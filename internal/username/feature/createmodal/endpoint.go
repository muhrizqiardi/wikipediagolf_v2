package createmodal

import (
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

type endpointDeps struct {
	Service     Service
	Template    *template.Template
	AuthContext authcontext.AuthContext
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /usernames/check", handler(deps.Service, deps.Template, deps.AuthContext))
}
