package joinroompage

import (
	"html/template"
	"net/http"
)

type EndpointDeps struct {
	Template *template.Template
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("/rooms/join", Handler(deps.Template))
}
