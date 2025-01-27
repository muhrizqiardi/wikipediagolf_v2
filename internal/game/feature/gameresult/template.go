package gameresult

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/result-page.html")
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "result-page.html", nil)
}
