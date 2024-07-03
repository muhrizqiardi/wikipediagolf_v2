package waitingpage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			path     = "/rooms"
			res      = httptest.NewRecorder()
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			serveMux = http.NewServeMux()
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		deps := EndpointDeps{
			Template: tmpl,
		}
		AddEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertEqual(t, res.Result().StatusCode, http.StatusOK)
	})
}
