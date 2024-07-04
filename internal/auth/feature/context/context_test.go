package context

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestSetRequest(t *testing.T) {
	var (
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		c   = NewAuthContext()
	)

	c.SetRequest(req, Val{
		UID: "testUID",
	})

	testutil.AssertEqualCMP(t, Val{UID: "testUID"}, req.Context().Value(key))
}

func TestGetFromRequest(t *testing.T) {
	var (
		r = httptest.NewRequest(http.MethodGet, "/", nil)
		c = NewAuthContext()
	)

	c.SetRequest(r, Val{
		UID: "testUID",
	})
	v, ok := c.GetFromRequest(r)
	testutil.AssertEqual(t, true, ok)
	testutil.AssertEqualCMP(t, Val{UID: "testUID"}, v)
}
