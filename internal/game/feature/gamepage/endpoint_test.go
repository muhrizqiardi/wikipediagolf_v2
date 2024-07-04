package gamepage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestEndpoint(t *testing.T) {
	t.Run("should not respond with 404", func(t *testing.T) {
		var (
			path     = "/game"
			res      = httptest.NewRecorder()
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			serveMux = http.NewServeMux()
		)
		tmpl := template.New("")
		tmpl, err := addTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)
		deps := EndpointDeps{
			Template: tmpl,
		}
		addEndpoint(serveMux, deps)

		serveMux.ServeHTTP(res, req)

		testutil.AssertInequal(t, res.Result().StatusCode, http.StatusNotFound)
	})
}