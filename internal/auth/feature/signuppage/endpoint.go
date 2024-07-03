package signuppage

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/sign-up", handler(deps.Template))
}
