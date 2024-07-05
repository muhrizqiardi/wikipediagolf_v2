package nicknamedialog

import (
	"html/template"
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

func Register(tmpl *template.Template, serveMux *http.ServeMux, actx authcontext.AuthContext) {
	addTemplate(tmpl)
	deps := endpointDeps{
		Template:    tmpl,
		AuthContext: actx,
	}
	addEndpoint(serveMux, deps)
}
