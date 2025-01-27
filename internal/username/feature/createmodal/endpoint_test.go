package createmodal

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	var (
		ms = &mockService{
			findByUIDV: &FindByUIDResponse{
				UID:      "mockUID",
				Username: "mockUsername",
			},
			findByUIDErr: nil,
		}
		body     = strings.NewReader(url.Values{"uid": []string{"mockUID"}}.Encode())
		res      = httptest.NewRecorder()
		req      = httptest.NewRequest(http.MethodPost, "/usernames/check", body)
		tmpl     = template.New("")
		serveMux = http.NewServeMux()
		c        = authcontext.NewAuthContext()
		deps     = endpointDeps{
			Service:     ms,
			Template:    tmpl,
			AuthContext: c,
		}
	)

	c.SetRequest(req, authcontext.Val{UID: "mockUID"})
	addEndpoint(serveMux, deps)

	serveMux.ServeHTTP(res, req)

	testutil.AssertInequal(t, http.StatusNotFound, res.Result().StatusCode)
}
