package createroom

import (
	"bytes"
	"fmt"
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
		fmt.Println(buf.String())
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, buf.Len())
	})
}
