package signuppage

import "html/template"

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/sign-up.html"
	return tmpl.ParseFS(templateFS, templateName)
}
