package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewServerDetailResponse is a constructor for the DetailResponse struct for a server.
func NewServerDetailResponse(currentUser *model.User, server *model.Server) *DetailResponse {
	headerText := "Server Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("servers.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/server/update/%d", server.ID)))
	}
	if currentUser.HasPrivilege("servers.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/server/delete/%d", server.ID)))
	}
	if currentUser.HasPrivilege("servers.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/server/list"))
	}
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

// ServerFormResponse is the struct for the server form responses.
type ServerFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewServerFormResponse is a constructor for the ServerFormResponse struct.
func NewServerFormResponse(title string, currentUser *model.User, server *model.Server, pools *model.Pools, runtimes *model.Runtimes) *ServerFormResponse {
	serverDetailResponse := NewServerDetailResponse(currentUser, server)
	serverDetailResponse.Header.Title = title
	serverDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	serverDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/server/list")}
	selectedPools := SelectedOptions{}
	selectedRuntimes := SelectedOptions{}
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
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: server.Name, Required: true},
		// Remote address.
		{Label: "Remote Address", Type: "text", Name: "remote_address", Value: server.RemoteAddr, Required: true},
		// Description.
		{Label: "Description", Type: "textarea", Name: "description", Value: server.Description, Required: false},
		// Pool.
		{Label: "Pool", Type: "checkboxgroup", Name: "pools", Options: pools.ToMap(), SelectedOptions: selectedPools},
		// Runtime.
		{Label: "Runtime", Type: "checkboxgroup", Name: "runtimes", Options: runtimes.ToMap(), SelectedOptions: selectedRuntimes},
	}
	return &ServerFormResponse{
		DetailResponse: serverDetailResponse,
		FormItems:      formItems,
	}
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
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", server.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: server.Name}}}
		columns = append(columns, nameColumn)
		remoteAddrColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: server.RemoteAddr}}}
		columns = append(columns, remoteAddrColumn)
		desctiptionColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: server.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
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
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
