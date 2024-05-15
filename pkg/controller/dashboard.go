package controller

import (
	"net/http"
)

// DashboardController is the dashboard controller.
func (c *Controller) DashboardController(w http.ResponseWriter, r *http.Request) {
	template := c.renderer.BuildTemplate("dashboard", []string{c.renderer.GetTemplateDirectoryPath() + "/dashboard/index.html.tmpl"})
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
