package signout

import "net/http"

func Register(serveMux *http.ServeMux) {
	addEndpoint(serveMux)
}
