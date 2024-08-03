package response

import (
	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/resources"
)

// Response is the struct for the response.
// It contains
// - the title of the page
// - the current user
// - the menu items for the side menu (based on the privileges)
// - the header block
type Response struct {
	Title       string
	CurrentUser *model.User
	SideMenu    []*components.Link
	Header      *components.ContentHeader
}

// NewResponse is a constructor for the Response struct.
func NewResponse(title string, user *model.User, header *components.ContentHeader) *Response {
	sideMenu := []*components.Link{}
	for _, resource := range resources.Resources {
		if user.HasPrivilege(resources.ResourcePrivileges[resource] + ".view") {
			sideMenu = append(sideMenu, components.NewLink(resource, "/admin/"+resource+"/list"))
		}
	}

	return &Response{
		Title:       title,
		CurrentUser: user,
		SideMenu:    sideMenu,
		Header:      header,
	}
}

// DetailResponse is the struct for the detail page.
// It contains the response and the details.
type DetailResponse struct {
	*Response
	Details *components.DetailItems
}

// NewDetailResponse is a constructor for the DetailResponse struct.
func NewDetailResponse(title string, currentUser *model.User, header *components.ContentHeader, details *components.DetailItems) *DetailResponse {
	return &DetailResponse{
		Response: NewResponse(title, currentUser, header),
		Details:  details,
	}
}

// SelectedOptions is the struct for the selected options.
type SelectedOptions []int64

// IsSelected is a helper function to check if the option is selected.
func (so SelectedOptions) IsSelected(option int64) bool {
	for _, selected := range so {
		if selected == option {
			return true
		}
	}
	return false
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
	// selected options. for select / checkbox group.
	SelectedOptions SelectedOptions
}

// ListingHeader is the struct for the listing header.
// It contains the header elements of the listing.
type ListingHeader struct {
	Headers []string
}

// ListingRow is the struct for the listing item.
// It contains the values of the item.
type ListingRow struct {
	Columns *ListingColumns
}

// ListingRows is the struct for the listing items.
type ListingRows []*ListingRow

// ListingColumn is the struct for the listing column.
type ListingColumn struct {
	Values *ListingColumnValues
}

// ListingColumns is the struct for the listing columns.
type ListingColumns []*ListingColumn

// ListingColumnValue is the struct for a listing column entry (one column might contain multiple values).
// It contains the value of the column.
// Also contains the link if it is a link.
// On case of the Form set to true, the link is the action of a POST form.
type ListingColumnValue struct {
	Value string
	Link  string
	Form  bool
}

// ListingColumnValues is the struct for the listing column values.
type ListingColumnValues []*ListingColumnValue

// Listing is the struct for the listing response.
// It contains the header block and the list items.
type Listing struct {
	Header *ListingHeader
	Rows   *ListingRows
}
