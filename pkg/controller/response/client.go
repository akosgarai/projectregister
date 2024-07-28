package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ClientDetailResponse is the struct for the client detail page.
type ClientDetailResponse struct {
	*Response
	Header *HeaderBlock
	Client *model.Client
}

// NewClientDetailResponse is a constructor for the ClientDetailResponse struct.
func NewClientDetailResponse(currentUser *model.User, client *model.Client) *ClientDetailResponse {
	return &ClientDetailResponse{
		Response: NewResponse("Client Detail", currentUser),
		Header: &HeaderBlock{
			Title:       "Client Detail",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Edit",
					Link:      fmt.Sprintf("/admin/client/update/%d", client.ID),
					Privilege: "clients.update",
				},
				{
					Label:     "Delete",
					Link:      fmt.Sprintf("/admin/client/delete/%d", client.ID),
					Privilege: "clients.delete",
				},
			},
		},
		Client: client,
	}
}

// ClientFormResponse is the struct for the client form responses.
type ClientFormResponse struct {
	*ClientDetailResponse
}

// NewClientFormResponse is a constructor for the ClientFormResponse struct.
func NewClientFormResponse(title string, currentUser *model.User, client *model.Client) *ClientFormResponse {
	clientDetailResponse := NewClientDetailResponse(currentUser, client)
	clientDetailResponse.Header.Title = title
	clientDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	clientDetailResponse.Header.Buttons = []*ActionButton{}
	return &ClientFormResponse{
		ClientDetailResponse: clientDetailResponse,
	}
}

// ClientListResponse is the struct for the client list page.
type ClientListResponse struct {
	*Response
	Header  *HeaderBlock
	Clients []*model.Client
}

// NewClientListResponse is a constructor for the ClientListResponse struct.
func NewClientListResponse(currentUser *model.User, clients []*model.Client) *ClientListResponse {
	return &ClientListResponse{
		Response: NewResponse("Client List", currentUser),
		Header: &HeaderBlock{
			Title:       "Client List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/client/create",
					Privilege: "clients.create",
				},
			},
		},
		Clients: clients,
	}
}
