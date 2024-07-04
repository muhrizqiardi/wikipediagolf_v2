package create

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler_ErrForbidden(t *testing.T) {
	var (
		res = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodPost, "/rooms", nil)
		ms  = &mockService{}
		c   = authcontext.NewAuthContext()
	)

	// request context wasn't set on middleware, unauthenticated
	handler(ms, c).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusForbidden, res.Result().StatusCode)
}

func TestHandler_ErrDuplicateRoomCode(t *testing.T) {
	var (
		res = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodPost, "/rooms", nil)
		ms  = &mockService{
			v:   nil,
			err: ErrDuplicateRoomCode,
		}
		c       = authcontext.NewAuthContext()
		mockUID = "mockUID"
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})
	handler(ms, c).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusInternalServerError, res.Result().StatusCode)
}

func TestHandler_NoError(t *testing.T) {
	var (
		res = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodPost, "/rooms", nil)
		ms  = &mockService{
			v:   &CreateRoomResponse{},
			err: nil,
		}
		c       = authcontext.NewAuthContext()
		mockUID = "mockUID"
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})
	handler(ms, c).ServeHTTP(res, req)

	testutil.AssertEqual(t, "/rooms", res.Result().Header.Get("HX-Location"))
}
