package waitingpage

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestExecuteTemplate(t *testing.T) {
	t.Run("should execute template", func(t *testing.T) {
		tmpl := template.New("")
		tmpl, err := AddTemplate(tmpl)
		testutil.AssertNoError(t, err)
		testutil.AssertNotNil(t, tmpl)
		var buf bytes.Buffer
		err = ExecuteTemplate(tmpl, &buf, TemplateData{
			Room: model.Room{
				ID:     [16]byte{},
				Code:   "",
				Status: "",
			},
			Members: []model.RoomMember{
				{
					ID:      uuid.New(),
					IsOwner: true,
					RoomID:  uuid.New(),
					UserUID: "123456",
				},
			},
		})
		testutil.AssertNoError(t, err)
		testutil.AssertInequal(t, 0, buf.Len())
	})
}
