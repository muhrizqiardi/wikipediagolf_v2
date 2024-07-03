package joinpage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should respond with create room page", func(t *testing.T) {
		var (
			path = "/rooms/join"
			res  = httptest.NewRecorder()
			req  = httptest.NewRequest(http.MethodGet, path, nil)
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)

		Handler(tmpl).ServeHTTP(res, req)

		testutil.AssertEqual(t, res.Result().StatusCode, http.StatusOK)
	})
}
