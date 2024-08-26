package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewServerDetailResponse is a constructor for the DetailResponse struct for a server.
func NewServerDetailResponse(currentUser *model.User, server *model.Server) *DetailResponse {
	headerText := "Server Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "servers", fmt.Sprintf("%d", server.ID)))
	runtimeValues := components.DetailValues{}
	if len(server.Runtimes) > 0 {
		for _, runtime := range server.Runtimes {
			runtimeValues = append(runtimeValues, &components.DetailValue{Value: runtime.Name, Link: fmt.Sprintf("/admin/runtime/view/%d", runtime.ID)})
		}
	}
	poolValues := components.DetailValues{}
	if len(server.Pools) > 0 {
		for _, pool := range server.Pools {
			poolValues = append(poolValues, &components.DetailValue{Value: pool.Name, Link: fmt.Sprintf("/admin/pool/view/%d", pool.ID)})
		}
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", server.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: server.Name}}},
		{Label: "Remote Address", Value: &components.DetailValues{{Value: server.RemoteAddr}}},
		{Label: "Description", Value: &components.DetailValues{{Value: server.Description}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: server.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: server.UpdatedAt}}},
		{Label: "Runtimes", Value: &runtimeValues},
		{Label: "Pools", Value: &poolValues},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateServerResponse is a constructor for the FormResponse struct for creating a new server.
func NewCreateServerResponse(currentUser *model.User, pools *model.Pools, runtimes *model.Runtimes) *FormResponse {
	return newServerFormResponse("Create Server", currentUser, &model.Server{}, pools, runtimes, "/admin/server/create", "POST", "Create")
}

// NewUpdateServerResponse is a constructor for the FormResponse struct for updating a server.
func NewUpdateServerResponse(currentUser *model.User, server *model.Server, pools *model.Pools, runtimes *model.Runtimes) *FormResponse {
	return newServerFormResponse("Update Server", currentUser, server, pools, runtimes, fmt.Sprintf("/admin/server/update/%d", server.ID), "POST", "Update")
}

// newServerFormResponse is a constructor for the FormResponse struct for a server.
func newServerFormResponse(title string, currentUser *model.User, server *model.Server, pools *model.Pools, runtimes *model.Runtimes, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/server/list")})
	var selectedPools, selectedRuntimes []int64
	if server.Pools != nil {
		for _, pool := range server.Pools {
			selectedPools = append(selectedPools, pool.ID)
		}
	}
	if server.Runtimes != nil {
		for _, runtime := range server.Runtimes {
			selectedRuntimes = append(selectedRuntimes, runtime.ID)
		}
	}
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", server.Name, true, nil, nil),
		// Remote address.
		components.NewFormItem("Remote Address", "remote_address", "text", server.RemoteAddr, true, nil, nil),
		// Description.
		components.NewFormItem("Description", "description", "textarea", server.Description, false, nil, nil),
		// Pool.
		components.NewFormItem("Pool", "pools", "checkboxgroup", "", true, pools.ToMap(), selectedPools),
		// Runtime.
		components.NewFormItem("Runtime", "runtimes", "checkboxgroup", "", true, runtimes.ToMap(), selectedRuntimes),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewServerListResponse is a constructor for the ListingResponse struct of the servers.
func NewServerListResponse(currentUser *model.User, servers *model.Servers) *ListingResponse {
	headerText := "Server List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("servers.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/server/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Remote Address", "Description", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("servers.update")
	userCanDelete := currentUser.HasPrivilege("servers.delete")
	for _, server := range *servers {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", server.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: server.Name}}}
		columns = append(columns, nameColumn)
		remoteAddrColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: server.RemoteAddr}}}
		columns = append(columns, remoteAddrColumn)
		desctiptionColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: server.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/server/view/%d", server.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/server/update/%d", server.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/server/delete/%d", server.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, nil)
}
