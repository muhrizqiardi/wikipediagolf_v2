package nicknamedialog

import (
	"html/template"
	"io"
)

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/choose-nickname-dialog.html"
	return tmpl.ParseFS(templateFS, templateName)
}

func executeTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "choose-nickname-dialog.html", nil)
}
