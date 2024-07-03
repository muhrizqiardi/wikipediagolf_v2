package createmodal

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	ctx "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
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
		deps     = EndpointDeps{
			Service:  ms,
			Template: tmpl,
		}
	)

	ctx.SetRequest(req, ctx.Val{UID: "mockUID"})
	AddEndpoint(serveMux, deps)

	serveMux.ServeHTTP(res, req)

	testutil.AssertInequal(t, http.StatusNotFound, res.Result().StatusCode)
}
