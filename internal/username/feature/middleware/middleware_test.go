package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestMiddleware_UsernameNotFound(t *testing.T) {
	var (
		ms = &mockService{
			findByUIDV:   nil,
			findByUIDErr: ErrUsernameNotFound,
		}
		nextCalled = 0
		next       = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { nextCalled++ })
		res        = httptest.NewRecorder()
		req        = httptest.NewRequest(http.MethodGet, "/", nil)
	)

	authcontext.SetRequest(req, authcontext.Val{UID: "mockUID"})

	Middleware(ms)(next).ServeHTTP(res, req)

	testutil.AssertInequal(t, 0, nextCalled)
	testutil.AssertEqual(t, http.StatusSeeOther, res.Result().StatusCode)
	testutil.AssertEqual(t, "/usernames/create", res.Result().Header.Get("Location"))
}

func TestMiddleware_UsernameFound(t *testing.T) {
	var (
		ms = &mockService{
			findByUIDV:   nil,
			findByUIDErr: ErrUsernameNotFound,
		}
		nextCalled = 0
		next       = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { nextCalled++ })
		res        = httptest.NewRecorder()
		req        = httptest.NewRequest(http.MethodGet, "/", nil)
	)

	authcontext.SetRequest(req, authcontext.Val{UID: "mockUID"})

	Middleware(ms)(next).ServeHTTP(res, req)

	testutil.AssertInequal(t, 0, nextCalled)
}
