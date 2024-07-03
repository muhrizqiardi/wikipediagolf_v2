package signin

import (
	"net/http"
)

type EndpointDeps struct {
	Service Service
}

func AddEndpoint(serveMux *http.ServeMux, deps EndpointDeps) {
	serveMux.Handle("POST /sign-in", Handler(deps.Service))
}
