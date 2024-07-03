package game

import (
	"html/template"
	"io"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/game.html")
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "game.html", nil)
}
