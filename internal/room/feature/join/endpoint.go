package join

import (
	"net/http"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
)

type endpointDeps struct {
	Service     Service
	AuthContext authcontext.AuthContext
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /rooms/join", handler(deps.Service, deps.AuthContext))
}
