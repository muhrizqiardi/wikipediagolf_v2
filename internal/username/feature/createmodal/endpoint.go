package createmodal

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Service  Service
	Template *template.Template
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("POST /usernames/check", Handler(deps.Service, deps.Template))
}
