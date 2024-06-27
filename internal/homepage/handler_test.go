package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestEndpoint(t *testing.T) {
	t.Run("should not return 404", func(t *testing.T) {
		var (
			req      = httptest.NewRequest(http.MethodGet, "/", nil)
			res      = httptest.NewRecorder()
			tmpl     = MustNewTemplate(NewTemplate())
			serveMux = http.NewServeMux()
			deps     = EndpointDeps{
				Template: tmpl,
			}
		)

		AddEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertInequal(t, 404, res.Result().StatusCode)
	})
}

func TestHandler(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			res = httptest.NewRecorder()
		)
		tmpl, err := NewTemplate()
		if err != nil {
			t.Error("exp nil; got err:", err)
		}

		Handler(tmpl).ServeHTTP(res, req)

		var (
			exp = 200
			got = res.Result().StatusCode
		)
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
}
