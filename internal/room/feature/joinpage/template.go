package joinpage

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/join-room.html")
}

func executeTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "join-room.html", nil)
}
