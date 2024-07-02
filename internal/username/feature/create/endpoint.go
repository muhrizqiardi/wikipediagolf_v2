package create

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
	Service  Service
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("POST /usernames/create", Handler(deps.Template, deps.Service))
}
