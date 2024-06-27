package homepage

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/{$}", Handler(deps.Template))
}
