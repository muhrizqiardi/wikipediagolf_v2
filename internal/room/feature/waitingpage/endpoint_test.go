package waitingpage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/partials"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddEndpoint(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			path     = "/rooms"
			res      = httptest.NewRecorder()
			req      = httptest.NewRequest(http.MethodGet, path, nil)
			serveMux = http.NewServeMux()
			c        = authcontext.NewAuthContext()
			ms       = &mockService{
				getRoomV: &GetRoomResponse{
					Room: model.Room{
						ID:     uuid.New(),
						Code:   "123456",
						Status: "open",
					},
					Members: []GetRoomResponseMember{},
				},
				getRoomErr: nil,
			}
			mockUID = "mockUID"
		)
		c.SetRequest(req, authcontext.Val{UID: mockUID})
		tmpl := template.New("")
		partials.Register(tmpl)
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		deps := EndpointDeps{
			Template:    tmpl,
			AuthContext: c,
			Service:     ms,
		}
		AddEndpoint(serveMux, deps)
		serveMux.ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
	})
}
