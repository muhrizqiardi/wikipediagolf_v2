package signup

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
			"email":           []string{"valid@email.com"},
			"username":        []string{"validUsername"},
			"password":        []string{"strongPassword?123"},
			"confirmPassword": []string{"strongPassword?123"},
		}.Encode())
		path        = "/sign-up"
		contentType = "application/x-www-form-urlencoded"
		res         = httptest.NewRecorder()
		req         = httptest.NewRequest(http.MethodGet, path, body)
		ms          = &mockService{
			v: &CreateUserResponse{
				UID:           "mockuid",
				Email:         "fulan@example.com",
				EmailVerified: false,
				PhoneNumber:   "123",
				DisplayName:   "mcok",
				PhotoURL:      "url",
				Disabled:      false,
			},
			err: nil,
		}
		serveMux = http.NewServeMux()
		deps     = endpointDeps{
			Service: ms,
		}
	)
	req.Header.Set("Content-Type", contentType)

	addEndpoint(serveMux, deps)
	serveMux.ServeHTTP(res, req)

	testutil.AssertInequal(t, http.StatusNotFound, res.Result().StatusCode)
}
