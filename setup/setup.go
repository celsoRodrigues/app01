package setup

import (
	"embed"
	"text/template"
)

func NewTemplates(templatesFs embed.FS) *template.Template {
	t := template.Must(template.New("").ParseFS(templatesFs, "templates/*.tpl"))
	return t
}
