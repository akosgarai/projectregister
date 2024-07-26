package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/model"
)

// DashboardController is the dashboard controller.
func (c *Controller) DashboardController(w http.ResponseWriter, r *http.Request) {
	content := struct {
		Title       string
		CurrentUser *model.User
	}{
		Title:       "Dashboard",
		CurrentUser: c.CurrentUser(r),
	}
	err := c.renderer.Template.RenderTemplate(w, "dashboard.html", content)
	if err != nil {
		panic(err)
	}
}
