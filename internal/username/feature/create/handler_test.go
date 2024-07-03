package create

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler_InvalidUsername(t *testing.T) {
	var (
		path = "/usernames/create"
		body = strings.NewReader(url.Values{
			"uid":      []string{"mockUID"},
			"username": []string{"invalid username"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		ms   = &mockService{createErr: ErrInvalidUsername}
		tmpl = template.New("")
	)

	addTemplate(tmpl)
	Handler(tmpl, ms).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Invalid username format", strings.TrimSpace(doc.Find(`[data-testid="createusernamealert"]`).First().Text()))
}

func TestHandler_DuplicateUsername(t *testing.T) {
	var (
		path = "/usernames/create"
		body = strings.NewReader(url.Values{
			"uid":      []string{"mockUID"},
			"username": []string{"username"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		ms   = &mockService{createErr: ErrDuplicateUsername}
		tmpl = template.New("")
	)

	addTemplate(tmpl)
	Handler(tmpl, ms).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, "Username already exists", strings.TrimSpace(doc.Find(`[data-testid="createusernamealert"]`).First().Text()))
}

func TestHandler_NoError(t *testing.T) {
	var (
		ms   = &mockService{createErr: nil}
		path = "/usernames/create"
		body = strings.NewReader(url.Values{
			"uid":      []string{"mockUID"},
			"username": []string{"validUsername"},
		}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, path, body)
		tmpl = template.New("")
	)

	addTemplate(tmpl)
	Handler(tmpl, ms).ServeHTTP(res, req)

	testutil.AssertEqual(t, res.Result().Header.Get("HX-Redirect"), "/")
	testutil.AssertEqual(t, res.Result().Header.Get("HX-Target"), "body")
	testutil.AssertEqual(t, res.Result().Header.Get("HX-Swap"), "outerHTML")
}
