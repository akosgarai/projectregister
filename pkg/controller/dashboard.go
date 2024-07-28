package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/controller/response"
)

// DashboardController is the dashboard controller.
func (c *Controller) DashboardController(w http.ResponseWriter, r *http.Request) {
	user := c.CurrentUser(r)
	header := &response.HeaderBlock{
		Title:       "Dashboard",
		CurrentUser: user,
		Buttons:     []*response.ActionButton{},
	}
	content := response.NewResponse("Dashboard", user, header)
	err := c.renderer.Template.RenderTemplate(w, "dashboard.html", content)
	if err != nil {
		panic(err)
	}
}
