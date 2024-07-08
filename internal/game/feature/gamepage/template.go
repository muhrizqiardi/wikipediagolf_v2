package gamepage

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/game.html")
}

type templateData struct {
	FromTitle        string
	FromTitleDecoded string
	ToTitle          string
	ToTitleDecoded   string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data templateData) error {
	return tmpl.ExecuteTemplate(wr, "game.html", data)
}
