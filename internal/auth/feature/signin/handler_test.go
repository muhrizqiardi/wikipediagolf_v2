package signin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler_NoError(t *testing.T) {
	var (
		body = strings.NewReader(url.Values{
			"idToken": []string{"mockIDToken"},
		}.Encode())
		path        = "/sign-in"
		contentType = "application/x-www-form-urlencoded"
		res         = httptest.NewRecorder()
		req         = httptest.NewRequest(http.MethodPost, path, body)
		ms          = &mockService{
			signInV: &SignInResponse{
				SessionCookie: "",
			},
			signInErr: nil,
		}
	)

	req.Header.Set("Content-Type", contentType)
	Handler(ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
}
