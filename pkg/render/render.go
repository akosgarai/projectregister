package render

import (
	"html/template"
)

const (
	// BaseTemplatePath is the base template.
	BaseTemplatePath = "web/template/base.html.tmpl"
)

// BuildTemplate builds a template.
func BuildTemplate(name string, files []string) *template.Template {
	return template.Must(template.New(name).ParseFiles(append(files, BaseTemplatePath)...))
}
