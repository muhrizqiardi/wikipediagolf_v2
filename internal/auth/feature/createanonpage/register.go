package createanonpage

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux) {
	addTemplate(tmpl)
	deps := endpointDeps{
		Template: tmpl,
	}
	addEndpoint(serveMux, deps)
}
