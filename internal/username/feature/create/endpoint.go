package create

import (
	"html/template"
	"net/http"
)

type endpointDeps struct {
	Template *template.Template
	Service  Service
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /usernames/create", handler(deps.Template, deps.Service))
}
