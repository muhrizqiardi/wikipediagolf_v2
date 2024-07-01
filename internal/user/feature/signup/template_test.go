package signup

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestAddTemplate(t *testing.T) {
	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	testutil.AssertNotNil(t, tmpl.Lookup("signup-alert-partial.html"))
}

func TestExecuteTemplate(t *testing.T) {
	tmpl := template.New("")
	tmpl, err := AddTemplate(tmpl)
	testutil.AssertNoError(t, err)
	var buf bytes.Buffer
	testutil.AssertNoError(t, ExecuteTemplate(tmpl, &buf, TemplateData{
		Message: "Invalid username",
	}))
}
