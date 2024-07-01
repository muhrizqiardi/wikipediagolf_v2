package asset

import "net/http"

func AssetHandler() http.Handler {
	return http.FileServerFS(assetFS)
}

func DistHandler() http.Handler {
	return http.FileServerFS(distFS)
}
