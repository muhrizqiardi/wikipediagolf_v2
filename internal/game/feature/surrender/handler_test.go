package surrender

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should not respond with 404", func(t *testing.T) {
		var (
			path = "/game/surrendered"
			res  = httptest.NewRecorder()
			req  = httptest.NewRequest(http.MethodGet, path, nil)
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)

		handler(tmpl).ServeHTTP(res, req)

		testutil.AssertInequal(t, res.Result().StatusCode, http.StatusNotFound)
	})
}
