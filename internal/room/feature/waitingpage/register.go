package waitingpage

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux) {
	AddTemplate(tmpl)
	deps := EndpointDeps{
		Template: tmpl,
	}
	AddEndpoint(serveMux, deps)
}
