package signup

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestHandler_InvalidUsername(t *testing.T) {
	var (
		ms = &mockService{
			v:   nil,
			err: ErrInvalidUsername,
		}
		path = "/sign-up"
		body = strings.NewReader(url.Values{
			"username":        []string{"invalid/username"},
			"email":           []string{"valid@example.com"},
			"password":        []string{"strongPassword?123"},
			"confirmPassword": []string{"strongPassword?123"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	Handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Invalid username format", strings.TrimSpace(doc.Find(`[data-testid="signupalert"]`).First().Text()))
}

func TestHandler_InvalidEmail(t *testing.T) {
	var (
		ms = &mockService{
			v:   nil,
			err: ErrInvalidEmail,
		}
		path = "/sign-up"
		body = strings.NewReader(url.Values{
			"username":        []string{"validUsername"},
			"email":           []string{"invalid example.com"},
			"password":        []string{"strongPassword?123"},
			"confirmPassword": []string{"strongPassword?123"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	Handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Invalid email format", strings.TrimSpace(doc.Find(`[data-testid="signupalert"]`).First().Text()))
}

func TestHandler_InvalidPassword(t *testing.T) {
	var (
		ms = &mockService{
			v:   nil,
			err: ErrInvalidPassword,
		}
		path = "/sign-up"
		body = strings.NewReader(url.Values{
			"username":        []string{"validUsername"},
			"email":           []string{"valid@example.com"},
			"password":        []string{"invld"},
			"confirmPassword": []string{"strongPassword?123"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	Handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Invalid password format", strings.TrimSpace(doc.Find(`[data-testid="signupalert"]`).First().Text()))
}

func TestHandler_ConfirmPasswordNotMatched(t *testing.T) {
	var (
		ms = &mockService{
			v:   nil,
			err: ErrPasswordNotMatch,
		}
		path = "/sign-up"
		body = strings.NewReader(url.Values{
			"username":        []string{"validUsername"},
			"email":           []string{"valid@example.com"},
			"password":        []string{"strongPassword?123"},
			"confirmPassword": []string{"inequal"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	Handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Confirm Password should be the same as Password", strings.TrimSpace(doc.Find(`[data-testid="signupalert"]`).First().Text()))
}

func TestHandler_NoError(t *testing.T) {
	var (
		ms = &mockService{
			v: &CreateUserResponse{
				Email: "someone@example.com",
			},
			err: nil,
		}
		path = "/sign-up"
		body = testutil.MustToFormUrlencoded(&CreateUserRequest{
			Email:    "someone@example.com",
			Password: "strong-Password123",
		})
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	Handler(ms, tmpl).ServeHTTP(res, req)

	testutil.AssertEqual(t, http.StatusSeeOther, res.Result().StatusCode)
}
