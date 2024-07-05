package waitingpage

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler_ErrForbidden(t *testing.T) {
	var (
		path = "/rooms"
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodGet, path, nil)
		c    = authcontext.NewAuthContext()
		ms   = &mockService{}
	)

	// no UID passed from middleware
	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)

	handler(tmpl, c, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusForbidden, res.Result().StatusCode)
}

func TestHandler_ErrService(t *testing.T) {
	var (
		path = "/rooms"
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodGet, path, nil)
		c    = authcontext.NewAuthContext()
		ms   = &mockService{
			getRoomV:   nil,
			getRoomErr: errors.New(""),
		}
		mockUID = "mockUID"
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})
	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)

	handler(tmpl, c, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusInternalServerError, res.Result().StatusCode)
}

func TestHandler_OK(t *testing.T) {
	var (
		path    = "/rooms"
		res     = httptest.NewRecorder()
		req     = httptest.NewRequest(http.MethodGet, path, nil)
		c       = authcontext.NewAuthContext()
		mockUID = "mockUID"
		ms      = &mockService{
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
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})

	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)

	handler(tmpl, c, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusOK, res.Result().StatusCode)
}
