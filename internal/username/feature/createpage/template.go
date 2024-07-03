package createpage

import (
	"html/template"
	"io"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/create-username-page.html")
}

type TemplateData struct {
	UID string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "create-username-page.html", nil)
}
