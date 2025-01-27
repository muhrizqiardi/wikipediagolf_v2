package asset

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler(t *testing.T) {
	t.Run("should return styles.css", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/dist/styles.css", nil)
			res = httptest.NewRecorder()
		)

		distHandler().ServeHTTP(res, req)

		var bodyBuf bytes.Buffer
		written, err := io.Copy(&bodyBuf, res.Result().Body)
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, written)
		testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
	})
}
