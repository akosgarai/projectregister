package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// EnvironmentDetailResponse is the struct for the environment detail page.
type EnvironmentDetailResponse struct {
	*Response
	Environment *model.Environment
}

// NewEnvironmentDetailResponse is a constructor for the EnvironmentDetailResponse struct.
func NewEnvironmentDetailResponse(currentUser *model.User, environment *model.Environment) *EnvironmentDetailResponse {
	header := &HeaderBlock{
		Title:       "Environment Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/environment/update/%d", environment.ID),
				Privilege: "environments.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/environment/delete/%d", environment.ID),
				Privilege: "environments.delete",
			},
		},
	}
	return &EnvironmentDetailResponse{
		Response:    NewResponse("Environment Detail", currentUser, header),
		Environment: environment,
	}
}

// EnvironmentFormResponse is the struct for the environment form responses.
type EnvironmentFormResponse struct {
	*EnvironmentDetailResponse
	FormItems []*FormItem
}

// NewEnvironmentFormResponse is a constructor for the EnvironmentFormResponse struct.
func NewEnvironmentFormResponse(title string, currentUser *model.User, environment *model.Environment, servers *model.Servers, databases *model.Databases) *EnvironmentFormResponse {
	environmentDetailResponse := NewEnvironmentDetailResponse(currentUser, environment)
	environmentDetailResponse.Header.Title = title
	environmentDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	environmentDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/environment/list"), Privilege: "environments.view"}}
	selectedServers := SelectedOptions{}
	selectedDatabases := SelectedOptions{}
	if environment.Servers != nil {
		for _, server := range environment.Servers {
			selectedServers = append(selectedServers, server.ID)
		}
	}
	if environment.Databases != nil {
		for _, database := range environment.Databases {
			selectedDatabases = append(selectedDatabases, database.ID)
		}
	}

	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: environment.Name, Required: true},
		// Description.
		{Label: "Description", Type: "textarea", Name: "description", Value: environment.Description},
		// Servers.
		{Label: "Servers", Type: "select", Name: "servers", Options: servers.ToMap(), SelectedOptions: selectedServers},
		// Databases.
		{Label: "Databases", Type: "select", Name: "databases", Options: databases.ToMap(), SelectedOptions: selectedDatabases},
	}
	return &EnvironmentFormResponse{
		EnvironmentDetailResponse: environmentDetailResponse,
		FormItems:                 formItems,
	}
}

// EnvironmentListResponse is the struct for the environment list page.
type EnvironmentListResponse struct {
	*Response
	Environments *model.Environments
}

// NewEnvironmentListResponse is a constructor for the EnvironmentListResponse struct.
func NewEnvironmentListResponse(currentUser *model.User, environments *model.Environments) *EnvironmentListResponse {
	header := &HeaderBlock{
		Title:       "Environment List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/environment/create",
				Privilege: "environments.create",
			},
		},
	}
	return &EnvironmentListResponse{
		Response:     NewResponse("Environment List", currentUser, header),
		Environments: environments,
	}
}
