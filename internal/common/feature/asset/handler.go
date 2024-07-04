package asset

import "net/http"

func assetHandler() http.Handler {
	return http.FileServerFS(assetFS)
}

func distHandler() http.Handler {
	return http.FileServerFS(distFS)
}
