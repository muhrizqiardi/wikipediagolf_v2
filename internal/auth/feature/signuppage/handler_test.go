package signuppage

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should return page", func(t *testing.T) {
		var (
			path = "/sign-up"
			req  = httptest.NewRequest(http.MethodGet, path, nil)
			res  = httptest.NewRecorder()
		)
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)

		Handler(tmpl).ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
		var buf bytes.Buffer
		written, err := io.Copy(&buf, res.Result().Body)
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, written)
	})
}
