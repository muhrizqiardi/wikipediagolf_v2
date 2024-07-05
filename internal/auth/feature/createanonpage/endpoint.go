package createanonpage

import (
	"html/template"
	"net/http"
)

type endpointDeps struct {
	Template *template.Template
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("/user/anon", handler(deps.Template))
}
