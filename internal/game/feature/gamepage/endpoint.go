package gamepage

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/game", Handler(deps.Template))
}
