package render

import (
	"encoding/json"
	"html/template"
	"net/http"
)

const (
	// BaseTemplatePath is the base template.
	BaseTemplatePath = "web/template/base.html.tmpl"
)

// BuildTemplate builds a template.
func BuildTemplate(name string, files []string) *template.Template {
	return template.Must(template.New(name).ParseFiles(append(files, BaseTemplatePath)...))
}

// JSON renders a JSON response.
func JSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
