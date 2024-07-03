package createmodal

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	ctx "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/feature/context"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestHandler_UsernameFound(t *testing.T) {
	var (
		ms = &mockService{
			findByUIDV: &FindByUIDResponse{
				UID:      "mockUID",
				Username: "mockUsername",
			},
			findByUIDErr: nil,
		}
		body = strings.NewReader(url.Values{"uid": []string{"mockUID"}}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, "/usernames/check", body)
		tmpl = template.New("")
	)

	ctx.SetRequest(req, ctx.Val{UID: "mockUID"})

	addTemplate(tmpl)
	handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertEqual(t, 0, doc.Find("dialog").Length())
}

func TestHandler_UsernameNotFound(t *testing.T) {
	var (
		ms = &mockService{
			findByUIDErr: ErrUsernameNotFound,
		}
		body = strings.NewReader(url.Values{"uid": []string{"mockUID"}}.Encode())
		res  = httptest.NewRecorder()
		req  = httptest.NewRequest(http.MethodPost, "/usernames/check", body)
		tmpl = template.New("")
	)

	ctx.SetRequest(req, ctx.Val{UID: "mockUID"})

	addTemplate(tmpl)
	handler(ms, tmpl).ServeHTTP(res, req)

	doc, err := goquery.NewDocumentFromReader(res.Result().Body)
	testutil.AssertNoError(t, err)
	testutil.AssertInequal(t, 0, doc.Find("dialog").Length())
}
