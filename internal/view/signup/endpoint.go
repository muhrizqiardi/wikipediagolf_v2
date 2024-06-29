package signup

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/sign-up", Handler(deps.Template))
}
