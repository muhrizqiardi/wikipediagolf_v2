package signin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	var (
		body = strings.NewReader(url.Values{
			"email":    []string{"valid@email.com"},
			"password": []string{"strongPassword?123"},
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
		serveMux = http.NewServeMux()
		deps     = endpointDeps{
			service: ms,
		}
	)
	req.Header.Set("Content-Type", contentType)

	addEndpoint(serveMux, deps)
	serveMux.ServeHTTP(res, req)

	testutil.AssertInequal(t, http.StatusNotFound, res.Result().StatusCode)
}
