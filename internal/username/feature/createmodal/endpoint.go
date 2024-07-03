package createmodal

import (
	"html/template"
	"net/http"
)

type endpointDeps struct {
	Service  Service
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /usernames/check", handler(deps.Service, deps.Template))
}
