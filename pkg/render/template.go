package render

import (
	"html/template"
	"net/http"
)

// TemplateInterface is an interface for Templates.
type TemplateInterface interface {
	SetBaseTemplate(baseTemplate string)
	AddTemplate(name string, files []string)
	RenderTemplate(w http.ResponseWriter, name string, data interface{}) error
}

// Templates is a struct that holds the templates that we want to use.
type Templates struct {
	baseTemplate string
	templates    map[string]*template.Template
}

// NewTemplates creates a new Templates struct.
func NewTemplates() *Templates {
	return &Templates{
		baseTemplate: "",
		templates:    make(map[string]*template.Template),
	}
}

// SetBaseTemplate sets the base template.
func (t *Templates) SetBaseTemplate(baseTemplate string) {
	t.baseTemplate = baseTemplate
}

// AddTemplate adds the templates to the Templates struct.
func (t *Templates) AddTemplate(name string, files []string) {
	t.templates[name] = template.Must(template.New(name).ParseFiles(append(files, t.baseTemplate)...))
}

// RenderTemplate renders the template.
func (t *Templates) RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates[name].ExecuteTemplate(w, "base.html", data)
}
