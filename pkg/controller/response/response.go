package response

import (
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/resources"
)

// SideMenuItem is the struct for the side menu items.
// It contains
// - the label of the item
// - the link.
type SideMenuItem struct {
	Label string
	Link  string
}

// Response is the struct for the response.
// It contains
// - the title of the page
// - the current user
// - the menu items for the side menu (based on the privileges)
type Response struct {
	Title       string
	CurrentUser *model.User
	SideMenu    []*SideMenuItem
}

// NewResponse is a constructor for the Response struct.
func NewResponse(title string, user *model.User) *Response {
	sideMenu := []*SideMenuItem{}
	for _, resource := range resources.Resources {
		if user.HasPrivilege(resources.ResourcePrivileges[resource] + ".view") {
			sideMenu = append(sideMenu, &SideMenuItem{
				Label: resource,
				Link:  "/admin/" + resource + "/list",
			})
		}
	}

	return &Response{
		Title:       title,
		CurrentUser: user,
		SideMenu:    sideMenu,
	}
}

// ActionButton is the struct for the action buttons.
// It contains the label of the button and the link.
// Also contains the necessary privileges to see the button.
type ActionButton struct {
	Label     string
	Link      string
	Privilege string
}

// HeaderBlock is the struct for the header block.
// It contains the title of the page.
// Also contains the action buttons.
type HeaderBlock struct {
	Title       string
	CurrentUser *model.User
	Buttons     []*ActionButton
}
