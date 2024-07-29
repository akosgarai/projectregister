package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ServerDetailResponse is the struct for the server detail page.
type ServerDetailResponse struct {
	*Response
	Server *model.Server
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
		},
	}
	return &ServerDetailResponse{
		Response: NewResponse("Server Detail", currentUser, header),
		Server:   server,
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
	Servers *model.Servers
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
	return &ServerListResponse{
		Response: NewResponse("Server List", currentUser, header),
		Servers:  servers,
	}
}
