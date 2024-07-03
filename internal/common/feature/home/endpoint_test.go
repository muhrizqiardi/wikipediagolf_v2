package home

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestEndpoint(t *testing.T) {
	t.Run("should not return 404", func(t *testing.T) {
		var (
			req      = httptest.NewRequest(http.MethodGet, "/", nil)
			res      = httptest.NewRecorder()
			serveMux = http.NewServeMux()
		)

		tmpl := template.New("")
		tmpl, err := addTemplate(tmpl)
		testutil.AssertNoError(t, err)
		deps := endpointDeps{
			Template: tmpl,
		}

		addEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertInequal(t, 404, res.Result().StatusCode)
	})
}
