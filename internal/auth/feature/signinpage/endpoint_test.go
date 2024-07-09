package signinpage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/partials"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestEndpoint(t *testing.T) {
	t.Run("should not return 404", func(t *testing.T) {
		var (
			path     = "/sign-in"
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			res      = httptest.NewRecorder()
			serveMux = http.NewServeMux()
		)

		tmpl := template.New("")
		partials.Register(tmpl)
		testutil.AssertNotNil(t, tmpl)
		tmpl, err := addTemplate(tmpl)
		if err != nil {
			t.Error("exp nil; got err:", err)
		}
		deps := EndpointDeps{
			Template: tmpl,
		}

		addEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertInequal(t, 404, res.Result().StatusCode)
	})
}
