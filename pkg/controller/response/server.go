package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ServerDetailResponse is the struct for the server detail page.
type ServerDetailResponse struct {
	*Response
	Header *HeaderBlock
	Server *model.Server
}

// NewServerDetailResponse is a constructor for the ServerDetailResponse struct.
func NewServerDetailResponse(currentUser *model.User, server *model.Server) *ServerDetailResponse {
	return &ServerDetailResponse{
		Response: NewResponse("Server Detail", currentUser),
		Header: &HeaderBlock{
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
		},
		Server: server,
	}
}

// ServerFormResponse is the struct for the server form responses.
type ServerFormResponse struct {
	*ServerDetailResponse
	Pools    []*model.Pool
	Runtimes []*model.Runtime
}

// NewServerFormResponse is a constructor for the ServerFormResponse struct.
func NewServerFormResponse(title string, currentUser *model.User, server *model.Server, pools []*model.Pool, runtimes []*model.Runtime) *ServerFormResponse {
	serverDetailResponse := NewServerDetailResponse(currentUser, server)
	serverDetailResponse.Header.Title = title
	serverDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	serverDetailResponse.Header.Buttons = []*ActionButton{}
	return &ServerFormResponse{
		ServerDetailResponse: serverDetailResponse,
		Pools:                pools,
		Runtimes:             runtimes,
	}
}

// ServerListResponse is the struct for the server list page.
type ServerListResponse struct {
	*Response
	Header  *HeaderBlock
	Servers []*model.Server
}

// NewServerListResponse is a constructor for the ServerListResponse struct.
func NewServerListResponse(currentUser *model.User, servers []*model.Server) *ServerListResponse {
	return &ServerListResponse{
		Response: NewResponse("Server List", currentUser),
		Header: &HeaderBlock{
			Title:       "Server List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/server/create",
					Privilege: "servers.create",
				},
			},
		},
		Servers: servers,
	}
}
