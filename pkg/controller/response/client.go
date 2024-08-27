package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewClientDetailResponse is a constructor for the DetailResponse struct for a client.
func NewClientDetailResponse(currentUser *model.User, client *model.Client) *DetailResponse {
	headerText := "Client Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "clients", fmt.Sprintf("%d", client.ID)))
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", client.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: client.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: client.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: client.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// newClientFormResponse is a constructor for the ClientFormResponse struct.
func newClientFormResponse(title string, currentUser *model.User, client *model.Client, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/client/list")})

	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", client.Name, true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}

	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewCreateClientResponse is a constructor for the FormResponse struct for the client create page.
func NewCreateClientResponse(currentUser *model.User) *FormResponse {
	return newClientFormResponse("Create Client", currentUser, &model.Client{}, "/admin/client/create", "POST", "Create")
}

// NewUpdateClientResponse is a constructor for the FormResponse struct for the client update page.
func NewUpdateClientResponse(currentUser *model.User, client *model.Client) *FormResponse {
	return newClientFormResponse("Update Client", currentUser, client, fmt.Sprintf("/admin/client/update/%d", client.ID), "POST", "Update")
}

// NewClientListResponse is a constructor for the ListingResponse struct of the clients.
func NewClientListResponse(currentUser *model.User, clients *model.Clients, filter *model.ClientFilter) *ListingResponse {
	headerText := "Client List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("clients.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/client/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("clients.update")
	userCanDelete := currentUser.HasPrivilege("clients.delete")
	for _, client := range *clients {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", client.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: client.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/client/view/%d", client.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/client/update/%d", client.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/client/delete/%d", client.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	/* Create the search form. The only form item is the name. */
	formItems := []*components.FormItem{
		components.NewFormItem("Name", "name", "text", filter.Name, false, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/client/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
