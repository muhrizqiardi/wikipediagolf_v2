package signup

import (
	"html/template"
	"net/http"
)

type endpointDeps struct {
	Service  Service
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /sign-up", handler(deps.Service, deps.Template))
}
