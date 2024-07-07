package partials

import "html/template"

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/*.html")
}
