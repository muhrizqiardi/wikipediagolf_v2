package signup

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/signup-alert-partial.html")
}

type TemplateData struct {
	// "error" or "success"
	Type    string
	Message string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "signup-alert-partial.html", data)
}
