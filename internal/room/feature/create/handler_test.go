package create

import (
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
		res     = httptest.NewRecorder()
		req     = httptest.NewRequest(http.MethodPost, "/rooms", nil)
		c       = authcontext.NewAuthContext()
		mockUID = "mockUID"
		ms      = &mockService{
			v: &CreateRoomResponse{
				Room: model.Room{
					ID:     uuid.New(),
					Code:   "123456",
					Status: "open",
				},
				Owner: model.RoomMember{
					ID:      uuid.New(),
					IsOwner: true,
					RoomID:  uuid.New(),
					UserUID: mockUID,
				},
			},
			err: nil,
		}
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})
	handler(ms, c).ServeHTTP(res, req)

	testutil.AssertEqual(t, "/rooms", res.Result().Header.Get("HX-Redirect"))
}
