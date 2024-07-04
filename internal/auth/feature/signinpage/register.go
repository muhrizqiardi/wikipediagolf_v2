package signinpage

import (
	"html/template"
	"net/http"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux) {
	addTemplate(tmpl)
	addEndpoint(serveMux, EndpointDeps{
		Template: tmpl,
	})
}
