package signup

import (
	"html/template"
	"io"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/signup-error-alert-partial.html")
}

type TemplateData struct {
	ErrorMessage string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "signup-error-alert-partial.html", data)
}
