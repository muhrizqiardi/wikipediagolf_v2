package nicknamedialog

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/choose-nickname-dialog.html"
	return tmpl.ParseFS(templateFS, templateName)
}

type templateData struct {
	// `join` or `create-room`
	Type string
}

func executeTemplate(tmpl *template.Template, wr io.Writer, data templateData) error {
	return tmpl.ExecuteTemplate(wr, "choose-nickname-dialog.html", data)
}
