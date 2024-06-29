package pregamesplashscreen

import (
	"html/template"
	"io"
)

func AddTemplate(tmpl *template.Template) (*template.Template, error) {
	return tmpl.ParseFS(templateFS, "template/pregame-splash-screen.html")
}

func ExecuteTemplate(tmpl *template.Template, wr io.Writer) error {
	return tmpl.ExecuteTemplate(wr, "pregame-splash-screen.html", nil)
}
