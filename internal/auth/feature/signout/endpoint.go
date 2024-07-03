package signout

import "net/http"

func addEndpoint(serveMux *http.ServeMux) {
	serveMux.Handle("DELETE /sign-out", handler())
}
