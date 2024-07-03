package createpage

import (
	"html/template"
	"net/http"
)

type endpointDeps struct {
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("/rooms/create", Handler(deps.Template))
}
