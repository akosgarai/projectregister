package render

import (
	"encoding/json"
	"net/http"

	"github.com/akosgarai/projectregister/pkg/config"
)

const (
	// BaseTemplatePath is the base template.
	BaseTemplatePath = "web/template/base.html.tmpl"
)

// Renderer is the renderer.
type Renderer struct {
	// templateDirectoryPath is the path to the template directory.
	templateDirectoryPath string

	// staticDirectoryPath is the path to the static directory.
	staticDirectoryPath string

	// the compiled templates
	Template TemplateInterface
}

// NewRenderer creates a new renderer.
func NewRenderer(envConfig *config.Environment, t TemplateInterface) *Renderer {
	t.SetBaseTemplate(envConfig.GetRenderTemplateDirectoryPath() + "/" + envConfig.GetRenderBaseTemplate())
	return &Renderer{
		templateDirectoryPath: envConfig.GetRenderTemplateDirectoryPath(),

		staticDirectoryPath: envConfig.GetStaticDirectoryPath(),

		Template: t,
	}
}

// GetTemplateDirectoryPath returns the template directory path.
func (r *Renderer) GetTemplateDirectoryPath() string {
	return r.templateDirectoryPath
}

// GetStaticDirectoryPath returns the static directory path.
func (r *Renderer) GetStaticDirectoryPath() string {
	return r.staticDirectoryPath
}

// JSON renders a JSON response.
func (r *Renderer) JSON(w http.ResponseWriter, status int, v interface{}) {
	// check that v is marshalable
	_, err := json.Marshal(v)
	if err != nil {
		r.Error(w, http.StatusInternalServerError, "Internal server error", err)
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
func (r *Renderer) Error(w http.ResponseWriter, status int, message string, details error) {
	if details != nil {
		message += " " + details.Error()
	}
	http.Error(w, message, status)
}
