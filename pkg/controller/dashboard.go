package controller

import (
	"net/http"

	"github.com/akosgarai/projectregister/pkg/controller/response"
	"github.com/akosgarai/projectregister/pkg/controller/response/components"
)

// DashboardController is the dashboard controller.
func (c *Controller) DashboardController(w http.ResponseWriter, r *http.Request) {
	user := c.CurrentUser(r)
	headerText := "Dashboard"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	content := response.NewResponse(headerText, user, headerContent)
	err := c.renderer.Template.RenderTemplate(w, "dashboard.html", content)
	if err != nil {
		panic(err)
	}
}
