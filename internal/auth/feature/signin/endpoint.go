package signin

import (
	"net/http"
)

type endpointDeps struct {
	service Service
}

func addEndpoint(serveMux *http.ServeMux, deps endpointDeps) {
	serveMux.Handle("POST /sign-in", handler(deps.service))
}
