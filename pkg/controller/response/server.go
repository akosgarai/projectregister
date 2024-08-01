package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ServerResponse is the struct for the server page.
type ServerResponse struct {
	*Response
	Server *model.Server
}

// ServerDetailResponse is the struct for the server detail page.
type ServerDetailResponse struct {
	*ServerResponse
	Details *DetailItems
}

// NewServerDetailResponse is a constructor for the ServerDetailResponse struct.
func NewServerDetailResponse(currentUser *model.User, server *model.Server) *ServerDetailResponse {
	header := &HeaderBlock{
		Title:       "Server Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/server/update/%d", server.ID),
				Privilege: "servers.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/server/delete/%d", server.ID),
				Privilege: "servers.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/server/list",
				Privilege: "servers.view",
			},
		},
	}
	runtimeValues := DetailValues{}
	if len(server.Runtimes) > 0 {
		for _, runtime := range server.Runtimes {
			runtimeValues = append(runtimeValues, &DetailValue{Value: runtime.Name, Link: fmt.Sprintf("/admin/runtime/view/%d", runtime.ID)})
		}
	}
	poolValues := DetailValues{}
	if len(server.Pools) > 0 {
		for _, pool := range server.Pools {
			poolValues = append(poolValues, &DetailValue{Value: pool.Name, Link: fmt.Sprintf("/admin/pool/view/%d", pool.ID)})
		}
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", server.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: server.Name}}},
		{Label: "Remote Address", Value: &DetailValues{{Value: server.RemoteAddr}}},
		{Label: "Description", Value: &DetailValues{{Value: server.Description}}},
		{Label: "Created At", Value: &DetailValues{{Value: server.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: server.UpdatedAt}}},
		{Label: "Runtimes", Value: &runtimeValues},
		{Label: "Pools", Value: &poolValues},
	}
	return &ServerDetailResponse{
		ServerResponse: &ServerResponse{
			Response: NewResponse("Server Detail", currentUser, header),
			Server:   server,
		},
		Details: details,
	}
}

// ServerFormResponse is the struct for the server form responses.
type ServerFormResponse struct {
	*ServerDetailResponse
	FormItems []*FormItem
}

// NewServerFormResponse is a constructor for the ServerFormResponse struct.
func NewServerFormResponse(title string, currentUser *model.User, server *model.Server, pools *model.Pools, runtimes *model.Runtimes) *ServerFormResponse {
	serverDetailResponse := NewServerDetailResponse(currentUser, server)
	serverDetailResponse.Header.Title = title
	serverDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	serverDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/server/list"), Privilege: "servers.view"}}
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
		ServerDetailResponse: serverDetailResponse,
		FormItems:            formItems,
	}
}

// ServerListResponse is the struct for the server list page.
type ServerListResponse struct {
	*Response
	Listing *Listing
}

// NewServerListResponse is a constructor for the ServerListResponse struct.
func NewServerListResponse(currentUser *model.User, servers *model.Servers) *ServerListResponse {
	header := &HeaderBlock{
		Title:       "Server List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/server/create",
				Privilege: "servers.create",
			},
		},
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Remote Address", "Description", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("servers.update")
	userCanDelete := currentUser.HasPrivilege("servers.delete")
	for _, server := range *servers {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", server.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: server.Name}}}
		columns = append(columns, nameColumn)
		remoteAddrColumn := &ListingColumn{&ListingColumnValues{{Value: server.RemoteAddr}}}
		columns = append(columns, remoteAddrColumn)
		desctiptionColumn := &ListingColumn{&ListingColumnValues{{Value: server.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/server/view/%d", server.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/server/update/%d", server.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/server/delete/%d", server.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &ServerListResponse{
		Response: NewResponse("Server List", currentUser, header),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
