package signup

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

type mockUserService struct {
}

func (s *mockUserService) Create(payload CreateUserRequest) (*CreateUserResponse, error) {
	return nil, nil
}

type mockUserService_errDuplicate struct {
}

func (s *mockUserService_errDuplicate) SignUp(payload CreateUserRequest) (*CreateUserResponse, error) {
	return nil, ErrDuplicateUser
}

func TestHandler(t *testing.T) {
	t.Run("should return error when user already exists", func(t *testing.T) {
		var (
			ms   = new(mockUserService_errDuplicate)
			path = "/users"

			body = testutil.MustToFormUrlencoded(&CreateUserRequest{
				Email:    "someone@example.com",
				Password: "strong-Password123",
			})

			res = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodPost, path, body)
		)

		Handler(ms).ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusConflict, res.Result().StatusCode)
	})
	t.Run("should return error when body is invalid", func(t *testing.T) {
		var (
			ms   = new(mockUserService_errDuplicate)
			path = "/sign-up"
			body = strings.NewReader("invalidbody")
			res  = httptest.NewRecorder()
			req  = httptest.NewRequest(http.MethodPost, path, body)
		)

		Handler(ms).ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusBadRequest, res.Result().StatusCode)
	})
	t.Run("should redirect to home", func(t *testing.T) {
		var (
			ms   = new(mockUserService_errDuplicate)
			path = "/sign-up"

			body = testutil.MustToFormUrlencoded(&CreateUserRequest{
				Email:    "someone@example.com",
				Password: "strong-Password123",
			})

			res = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodPost, path, body)
		)

		Handler(ms).ServeHTTP(res, req)

		testutil.AssertEqual(t, http.StatusConflict, res.Result().StatusCode)
	})
}
