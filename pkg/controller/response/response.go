package response

import (
	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/resources"
)

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
	SideMenu    []*components.Link
	Header      *HeaderBlock
}

// NewResponse is a constructor for the Response struct.
func NewResponse(title string, user *model.User, header *HeaderBlock) *Response {
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

// DetailValue is the struct for the detail values.
// It holds the value, and if it is a link, the link.
type DetailValue struct {
	Value string
	Link  string
}

// DetailValues is the struct for the detail values.
type DetailValues []*DetailValue

// DetailItem is the struct for the detail items.
// It holds the label and the value.
// The value is a list of DetailValue
type DetailItem struct {
	Label string
	Value *DetailValues
}

// DetailItems is the struct for the detail items.
type DetailItems []*DetailItem

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
