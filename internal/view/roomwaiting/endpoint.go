package roomwaiting

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/rooms", Handler(deps.Template))
}
