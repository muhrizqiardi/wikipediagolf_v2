package home

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			res = httptest.NewRecorder()
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)

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
