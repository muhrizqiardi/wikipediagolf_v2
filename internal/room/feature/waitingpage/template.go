package waitingpage

import (
	"html/template"
	"io"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/waiting-room.html")
}

type TemplateData struct {
	Room    model.Room
	Members []model.RoomMember
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "waiting-room.html", data)
}
