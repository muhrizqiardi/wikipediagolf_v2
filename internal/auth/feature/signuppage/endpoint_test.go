package signuppage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	t.Run("should respond with 200", func(t *testing.T) {
		var (
			path     = "/sign-up"
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			res      = httptest.NewRecorder()
			serveMux = http.NewServeMux()
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)
		deps := EndpointDeps{
			Template: tmpl,
		}

		AddEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
	})
}
