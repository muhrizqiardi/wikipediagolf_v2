package surrenderpage

import (
	"html/template"
	"io"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/surrender-partial.html")
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "surrender-partial.html", nil)
}
