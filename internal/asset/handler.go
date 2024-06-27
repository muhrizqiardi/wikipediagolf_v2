package asset

import "net/http"

func AddEndpoint() func(serveMux *http.ServeMux) {
	return func(serveMux *http.ServeMux) {
		serveMux.Handle("GET /assets/", http.FileServerFS(AssetFS))
	}
}
