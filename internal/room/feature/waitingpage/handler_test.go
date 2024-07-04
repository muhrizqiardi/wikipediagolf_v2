package waitingpage

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	authcontext "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
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

	testutil.AssertEqual(t, res.Result().StatusCode, http.StatusOK)
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
	)

	// no UID passed from middleware
	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)

	handler(tmpl, c, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, res.Result().StatusCode, http.StatusOK)
}

func TestHandler_OK(t *testing.T) {
	var (
		path    = "/rooms"
		res     = httptest.NewRecorder()
		req     = httptest.NewRequest(http.MethodGet, path, nil)
		c       = authcontext.NewAuthContext()
		mockUID = "mockUID"
		ms      = &mockService{
			getRoomV:   nil,
			getRoomErr: errors.New(""),
		}
	)

	c.SetRequest(req, authcontext.Val{UID: mockUID})

	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)

	handler(tmpl, c, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, res.Result().StatusCode, http.StatusOK)
}
