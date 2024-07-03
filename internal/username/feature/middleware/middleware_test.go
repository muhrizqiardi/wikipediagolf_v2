package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/authctx"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
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

	authctx.SetRequest(req, authctx.Val{UID: "mockUID"})

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

	authctx.SetRequest(req, authctx.Val{UID: "mockUID"})

	Middleware(ms)(next).ServeHTTP(res, req)

	testutil.AssertInequal(t, 0, nextCalled)
}
