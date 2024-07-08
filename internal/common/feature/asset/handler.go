package asset

import (
	"net/http"

	"github.com/muhrizqiardi/wikipediagolf_v2/client"
)

func distHandler() http.Handler {
	return http.FileServerFS(client.DistFS)
}
