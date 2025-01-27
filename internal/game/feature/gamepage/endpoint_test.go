package gamepage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/partials"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestEndpoint(t *testing.T) {
	t.Run("should not respond with 404", func(t *testing.T) {
		var (
			path     = "/game"
			res      = httptest.NewRecorder()
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			serveMux = http.NewServeMux()
			c        = authcontext.NewAuthContext()
		)
		c.SetRequest(req, authcontext.Val{
			UID:    "mockUID",
			IsAnon: false,
		})
		tmpl := template.New("")
		partials.Register(tmpl)
		tmpl, err := addTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)
		deps := EndpointDeps{
			Template:    tmpl,
			AuthContext: authcontext.NewAuthContext(),
			Service: &mockService{
				v:   &model.Game{},
				err: err,
			},
		}
		addEndpoint(serveMux, deps)

		serveMux.ServeHTTP(res, req)

		testutil.AssertInequal(t, res.Result().StatusCode, http.StatusNotFound)
	})
}
