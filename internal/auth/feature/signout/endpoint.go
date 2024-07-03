package signout

import "net/http"

func AddEndpoint(serveMux *http.ServeMux) {
	serveMux.Handle("DELETE /sign-out", Handler())
}
