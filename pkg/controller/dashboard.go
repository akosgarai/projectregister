package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/controller/response"
)

// DashboardController is the dashboard controller.
func (c *Controller) DashboardController(w http.ResponseWriter, r *http.Request) {
	content := response.NewResponse("Dashboard", c.CurrentUser(r))
	err := c.renderer.Template.RenderTemplate(w, "dashboard.html", content)
	if err != nil {
		panic(err)
	}
}
