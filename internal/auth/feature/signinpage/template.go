package signinpage

import "html/template"

func addTemplate(tmpl *template.Template) (*template.Template, error) {
	templateName := "template/sign-in.html"
	return tmpl.ParseFS(templateFS, templateName)
}
