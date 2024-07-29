package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ClientResponse is the struct for the client page.
type ClientResponse struct {
	*Response
	Client *model.Client
}

// ClientDetailResponse is the struct for the client detail page.
type ClientDetailResponse struct {
	*ClientResponse
	Details *DetailItems
}

// NewClientDetailResponse is a constructor for the ClientDetailResponse struct.
func NewClientDetailResponse(currentUser *model.User, client *model.Client) *ClientDetailResponse {
	header := &HeaderBlock{
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
			{
				Label:     "List",
				Link:      "/admin/client/list",
				Privilege: "clients.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", client.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: client.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: client.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: client.UpdatedAt}}},
	}
	return &ClientDetailResponse{
		ClientResponse: &ClientResponse{
			Response: NewResponse("Client Detail", currentUser, header),
			Client:   client,
		},
		Details: details,
	}
}

// ClientFormResponse is the struct for the client form responses.
type ClientFormResponse struct {
	*ClientDetailResponse
	FormItems []*FormItem
}

// NewClientFormResponse is a constructor for the ClientFormResponse struct.
func NewClientFormResponse(title string, currentUser *model.User, client *model.Client) *ClientFormResponse {
	clientDetailResponse := NewClientDetailResponse(currentUser, client)
	clientDetailResponse.Header.Title = title
	clientDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	clientDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/client/list"), Privilege: "clients.view"}}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: client.Name, Required: true},
	}
	return &ClientFormResponse{
		ClientDetailResponse: clientDetailResponse,
		FormItems:            formItems,
	}
}

// ClientListResponse is the struct for the client list page.
type ClientListResponse struct {
	*Response
	Clients *model.Clients
}

// NewClientListResponse is a constructor for the ClientListResponse struct.
func NewClientListResponse(currentUser *model.User, clients *model.Clients) *ClientListResponse {
	header := &HeaderBlock{
		Title:       "Client List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/client/create",
				Privilege: "clients.create",
			},
		},
	}
	return &ClientListResponse{
		Response: NewResponse("Client List", currentUser, header),
		Clients:  clients,
	}
}
