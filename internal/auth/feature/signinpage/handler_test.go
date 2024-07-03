package signinpage

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			res = httptest.NewRecorder()
		)
		tmpl := template.New("")
		tmpl, err := addTemplate(tmpl)
		if err != nil {
			t.Error("exp nil; got err:", err)
		}

		handler(tmpl).ServeHTTP(res, req)

		var (
			exp = 200
			got = res.Result().StatusCode
		)
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	})
}
