package render

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/akosgarai/projectregister/pkg/config"
)

const (
	// BaseTemplatePath is the base template.
	BaseTemplatePath = "web/template/base.html.tmpl"
)

// Renderer is the renderer.
type Renderer struct {
	// baseTemplate is the base template.
	baseTemplate string
}

// NewRenderer creates a new renderer.
func NewRenderer(envConfig *config.Environment) *Renderer {
	return &Renderer{
		baseTemplate: envConfig.GetRenderTemplateDirectoryPath() + "/" + envConfig.GetRenderBaseTemplate(),
	}
}

// BuildTemplate builds a template.
func (r *Renderer) BuildTemplate(name string, files []string) *template.Template {
	return template.Must(template.New(name).ParseFiles(append(files, r.baseTemplate)...))
}

// JSON renders a JSON response.
func (r *Renderer) JSON(w http.ResponseWriter, status int, v interface{}) {
	// check that v is marshalable
	_, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// Status renders a status response with empty body.
func (r *Renderer) Status(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// Error renders an error response.
func (r *Renderer) Error(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}
