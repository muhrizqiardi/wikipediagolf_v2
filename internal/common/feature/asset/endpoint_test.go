package asset

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	t.Run("should not return 404", func(t *testing.T) {
		{
			var (
				req      = httptest.NewRequest(http.MethodGet, "/dist/styles.css", nil)
				res      = httptest.NewRecorder()
				serveMux = http.NewServeMux()
			)

			AddEndpoint(serveMux)
			serveMux.ServeHTTP(res, req)

			testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
		}
	})
}
