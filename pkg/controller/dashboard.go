package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/render"
)

// DashboardController is the dashboard controller.
func DashboardController(w http.ResponseWriter, r *http.Request) {
	template := render.BuildTemplate("dashboard", []string{"web/template/dashboard/index.html.tmpl"})
	content := struct {
		Title string
	}{
		Title: "Dashboard",
	}
	err := template.ExecuteTemplate(w, "base.html", content)
	if err != nil {
		panic(err)
	}
}
