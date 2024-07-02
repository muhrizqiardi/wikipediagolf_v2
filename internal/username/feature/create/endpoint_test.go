package create

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestAddEndpoint(t *testing.T) {
	var (
		path = "/usernames/create"
		body = strings.NewReader(url.Values{
			"uid":      []string{"mockUID"},
			"username": []string{"username"},
		}.Encode())
		res      = httptest.NewRecorder()
		req      = httptest.NewRequest(http.MethodPost, path, body)
		serveMux = http.NewServeMux()
	)
	deps := EndpointDeps{
		Service: &mockService{
			createV: &CreateUsernameResponse{
				UID:       "",
				Username:  "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			createErr: nil,
		},
	}
	AddEndpoint(serveMux, deps)
	serveMux.ServeHTTP(res, req)

	testutil.AssertInequal(t, http.StatusNotFound, res.Result().StatusCode)
}
