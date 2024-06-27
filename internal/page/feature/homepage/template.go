package homepage

import (
	"embed"
	"html/template"
)

//go:embed template/*
var templateFS embed.FS

func NewTemplate() (*template.Template, error) {
	tmpl, err := template.ParseFS(templateFS, "template/*.html")
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func MustNewTemplate(tmpl *template.Template, err error) *template.Template {
	if err != nil {
		panic(1)
	}

	return tmpl
}
