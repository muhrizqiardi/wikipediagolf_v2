package create

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/create-username-alert-partial.html")
}

type TemplateData struct {
	// "error" or "success"
	Type    string
	Message string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "create-username-alert-partial.html", data)
}
