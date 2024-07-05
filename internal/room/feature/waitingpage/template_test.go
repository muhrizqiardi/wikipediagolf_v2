package waitingpage

import (
	"bytes"
	"html/template"
	"regexp"
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
			CurrentUserUID: "mockUIDCurrent",
			Room: model.Room{
				ID:     uuid.New(),
				Code:   "123456",
				Status: "open",
			},
			Members: []GetRoomResponseMember{
				{
					ID:          uuid.New(),
					IsOwner:     true,
					RoomID:      uuid.New(),
					UserUID:     "mockUID1",
					Username:    "Challenger123",
					DisplayName: "",
				},
				{
					ID:          uuid.New(),
					IsOwner:     false,
					RoomID:      uuid.New(),
					UserUID:     "mockUID2",
					Username:    "",
					DisplayName: "Kepler 22-b",
				},
			},
			MembersTotal: 3,
		})
		content := buf.String()
		testutil.AssertNoError(t, err)
		{
			ok, err := regexp.MatchString("@Challenger123", content)
			testutil.AssertNoError(t, err)
			if !ok {
				t.Error("exp match @Challenger123")
			}
		}
		{
			ok, err := regexp.MatchString("Kepler 22-b", content)
			testutil.AssertNoError(t, err)
			if !ok {
				t.Error("exp match Kepler 22-b")
			}
		}
		{
			ok, err := regexp.MatchString("123456", content)
			testutil.AssertNoError(t, err)
			if !ok {
				t.Error(`exp match 123456`)
			}
		}
		{
			ok, err := regexp.MatchString(`Player list \(3\)`, content)
			testutil.AssertNoError(t, err)
			if !ok {
				t.Error(`exp match Player list \(3\)`)
			}
		}
	})
}
