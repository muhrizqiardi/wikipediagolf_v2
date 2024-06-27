package asset

import "net/http"

func Handler() http.Handler {
	return http.FileServerFS(assetFS)
}
