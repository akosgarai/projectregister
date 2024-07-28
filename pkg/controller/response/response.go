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

// Response is the struct for the response.
// It contains
// - the title of the page
// - the current user
// - the menu items for the side menu (based on the privileges)
// - the header block
type Response struct {
	Title       string
	CurrentUser *model.User
	SideMenu    []*SideMenuItem
	Header      *HeaderBlock
}

// NewResponse is a constructor for the Response struct.
func NewResponse(title string, user *model.User, header *HeaderBlock) *Response {
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
		Header:      header,
	}
}

// FormItem is the struct for the form items.
type FormItem struct {
	Label    string
	Name     string
	Type     string
	Value    string
	Required bool
	// On case of select / checkbox group type we need the options.
	Options map[int64]string
}
