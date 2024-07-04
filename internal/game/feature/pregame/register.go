package pregame

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux) {
	addTemplate(tmpl)
	deps := EndpointDeps{
		Template: tmpl,
	}
	addEndpoint(serveMux, deps)
}
