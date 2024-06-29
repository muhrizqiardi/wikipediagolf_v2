package roomwaiting

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestExecuteTemplate(t *testing.T) {
	t.Run("should execute template", func(t *testing.T) {
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)
		var buf bytes.Buffer
		err = ExecuteTemplate(tmpl, &buf)
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, buf.Len())
	})
}