package asset

import (
	"net/http"

	"github.com/muhrizqiardi/wikipediagolf_v2/client"
)

func assetHandler() http.Handler {
	return http.FileServerFS(client.AssetFS)
}

func distHandler() http.Handler {
	return http.FileServerFS(client.DistFS)
}
