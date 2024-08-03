package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewClientDetailResponse is a constructor for the DetailResponse struct for a client.
func NewClientDetailResponse(currentUser *model.User, client *model.Client) *DetailResponse {
	headerText := "Client Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("clients.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/client/update/%d", client.ID)))
	}
	if currentUser.HasPrivilege("clients.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/client/delete/%d", client.ID)))
	}
	if currentUser.HasPrivilege("clients.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/client/list"))
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", client.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: client.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: client.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: client.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// ClientFormResponse is the struct for the client form responses.
type ClientFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewClientFormResponse is a constructor for the ClientFormResponse struct.
func NewClientFormResponse(title string, currentUser *model.User, client *model.Client) *ClientFormResponse {
	clientDetailResponse := NewClientDetailResponse(currentUser, client)
	clientDetailResponse.Header.Title = title
	clientDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	clientDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/client/list")}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: client.Name, Required: true},
	}
	return &ClientFormResponse{
		DetailResponse: clientDetailResponse,
		FormItems:      formItems,
	}
}

// ClientListResponse is the struct for the client list page.
type ClientListResponse struct {
	*Response
	Listing *Listing
}

// NewClientListResponse is a constructor for the ClientListResponse struct.
func NewClientListResponse(currentUser *model.User, clients *model.Clients) *ClientListResponse {
	headerText := "Client List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("clients.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/client/create"))
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("clients.update")
	userCanDelete := currentUser.HasPrivilege("clients.delete")
	for _, client := range *clients {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", client.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: client.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/client/view/%d", client.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/client/update/%d", client.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/client/delete/%d", client.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &ClientListResponse{
		Response: NewResponse(headerText, currentUser, headerContent),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
