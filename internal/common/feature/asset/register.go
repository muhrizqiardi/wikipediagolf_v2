package asset

import "net/http"

func Register(serveMux *http.ServeMux) {
	AddEndpoint(serveMux)
}
