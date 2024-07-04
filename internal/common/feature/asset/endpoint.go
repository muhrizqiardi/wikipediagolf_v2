package asset

import "net/http"

func AddEndpoint(serveMux *http.ServeMux) {
	serveMux.Handle("GET /dist/", distHandler())
	serveMux.Handle("GET /assets/", assetHandler())
}
