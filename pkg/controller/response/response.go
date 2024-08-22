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

// newDetailHeaderButtons is a helper function to generate the buttons for the detail page.
// The input resource is plural, as the plural is needed for the privilege check.
// The link need the singular version of the resource. It is the resource without the 's' at the end.
func newDetailHeaderButtons(currentUser *model.User, resource, id string) []*components.Link {
	buttons := []*components.Link{}
	resourceSingular := resource[:len(resource)-1]
	if currentUser.HasPrivilege(resource + ".update") {
		buttons = append(buttons, components.NewLink("Edit", "/admin/"+resourceSingular+"/update/"+id))
	}
	if currentUser.HasPrivilege(resource + ".delete") {
		buttons = append(buttons, components.NewLink("Delete", "/admin/"+resourceSingular+"/delete/"+id))
	}
	if currentUser.HasPrivilege(resource + ".view") {
		buttons = append(buttons, components.NewLink("List", "/admin/"+resourceSingular+"/list"))
	}
	return buttons
}

// ListingResponse is the struct for the listing page.
// It contains the response and the listing.
type ListingResponse struct {
	*Response
	Listing *components.Listing
}

// NewListingResponse is a constructor for the ListingResponse struct.
func NewListingResponse(title string, currentUser *model.User, header *components.ContentHeader, listing *components.Listing) *ListingResponse {
	return &ListingResponse{
		Response: NewResponse(title, currentUser, header),
		Listing:  listing,
	}
}

// FormResponse is the struct for the form page.
// It contains the response and the form items.
type FormResponse struct {
	*Response
	Form *components.Form
}

// NewFormResponse is a constructor for the FormResponse struct.
func NewFormResponse(title string, currentUser *model.User, header *components.ContentHeader, form *components.Form) *FormResponse {
	return &FormResponse{
		Response: NewResponse(title, currentUser, header),
		Form:     form,
	}
}

// ApplicationImportMappingResponse is the struct for the application import mapping page.
// It contains the preview listing and the mapping form.
type ApplicationImportMappingResponse struct {
	*Response
	Listing *components.Listing
	Form    *components.Form
}

// NewApplicationImportMappingResponse is a constructor for the ApplicationImportMappingResponse struct.
func NewApplicationImportMappingResponse(title string, currentUser *model.User, header *components.ContentHeader, previewListing *components.Listing, mappingForm *components.Form) *ApplicationImportMappingResponse {
	return &ApplicationImportMappingResponse{
		Response: NewResponse(title, currentUser, header),
		Listing:  previewListing,
		Form:     mappingForm,
	}
}
