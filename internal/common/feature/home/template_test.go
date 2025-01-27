package home

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/common/feature/partials"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestAddTemplate(t *testing.T) {
	t.Run("should add template to an existing template instance", func(t *testing.T) {
		templateName := "index.html"
		tmpl := template.New("")
		partials.Register(tmpl)

		tmpl, err := addTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)

		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, templateName, nil)
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, buf.Len())
	})
}
