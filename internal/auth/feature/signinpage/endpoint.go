package signinpage

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("GET /sign-in", handler(deps.Template))
}
