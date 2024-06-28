package homepage

import "html/template"

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/index.html"
	return tmpl.ParseFS(templateFS, templateName)
}
