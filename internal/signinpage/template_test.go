package signinpage

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestAddTemplate(t *testing.T) {
	t.Run("should add template to an existing template instance", func(t *testing.T) {
		templateName := "sign-in.html"
		tmpl := template.New("")

		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)

		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, templateName, nil)
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, buf.Len())
	})
}
