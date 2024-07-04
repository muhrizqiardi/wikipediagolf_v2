package createmodal

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/create-username-modal-partial.html")
}

type TemplateData struct {
	UID string
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer, data TemplateData) error {
	return tmpl.ExecuteTemplate(wr, "create-username-modal-partial.html", data)
}
