package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// EnvironmentDetailResponse is the struct for the environment detail page.
type EnvironmentDetailResponse struct {
	*Response
	Header      *HeaderBlock
	Environment *model.Environment
}

// NewEnvironmentDetailResponse is a constructor for the EnvironmentDetailResponse struct.
func NewEnvironmentDetailResponse(currentUser *model.User, environment *model.Environment) *EnvironmentDetailResponse {
	return &EnvironmentDetailResponse{
		Response: NewResponse("Environment Detail", currentUser),
		Header: &HeaderBlock{
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
		},
		Environment: environment,
	}
}

// EnvironmentFormResponse is the struct for the environment form responses.
type EnvironmentFormResponse struct {
	*EnvironmentDetailResponse
	Servers   []*model.Server
	Databases []*model.Database
}

// NewEnvironmentFormResponse is a constructor for the EnvironmentFormResponse struct.
func NewEnvironmentFormResponse(title string, currentUser *model.User, environment *model.Environment, servers []*model.Server, databases []*model.Database) *EnvironmentFormResponse {
	environmentDetailResponse := NewEnvironmentDetailResponse(currentUser, environment)
	environmentDetailResponse.Header.Title = title
	environmentDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	environmentDetailResponse.Header.Buttons = []*ActionButton{}
	return &EnvironmentFormResponse{
		EnvironmentDetailResponse: environmentDetailResponse,
		Servers:                   servers,
		Databases:                 databases,
	}
}

// EnvironmentListResponse is the struct for the environment list page.
type EnvironmentListResponse struct {
	*Response
	Header       *HeaderBlock
	Environments []*model.Environment
}

// NewEnvironmentListResponse is a constructor for the EnvironmentListResponse struct.
func NewEnvironmentListResponse(currentUser *model.User, environments []*model.Environment) *EnvironmentListResponse {
	return &EnvironmentListResponse{
		Response: NewResponse("Environment List", currentUser),
		Header: &HeaderBlock{
			Title:       "Environment List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/environment/create",
					Privilege: "environments.create",
				},
			},
		},
		Environments: environments,
	}
}
