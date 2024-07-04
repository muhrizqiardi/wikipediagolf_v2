package home

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/index.html"
	return tmpl.ParseFS(templateFS, templateName)
}

type TemplateData struct {
	IsAuthenticated bool
	UID             string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "index.html", data)
}
