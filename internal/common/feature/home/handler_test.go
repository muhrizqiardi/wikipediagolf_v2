package home

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			c   = authcontext.NewAuthContext()
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			res = httptest.NewRecorder()
		)
		tmpl := template.New("")
		tmpl, err := addTemplate(tmpl)
		testutil.AssertNoError(t, err)

		handler(tmpl, c).ServeHTTP(res, req)

		var (
			exp = 200
			got = res.Result().StatusCode
		)
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
}
