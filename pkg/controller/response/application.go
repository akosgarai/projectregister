package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationDetailResponse is the struct for the application view response.
// It extends the Response struct with a header block and the application.
type ApplicationDetailResponse struct {
	*Response
	Application *model.Application
}

// NewApplicationDetailResponse is a constructor for the ApplicationViewResponse struct.
func NewApplicationDetailResponse(user *model.User, app *model.Application) *ApplicationDetailResponse {
	header := &HeaderBlock{
		Title:       "Application View",
		CurrentUser: user,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/application/update/%d", app.ID),
				Privilege: "applications.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/application/delete/%d", app.ID),
				Privilege: "application.delete",
			},
		},
	}
	return &ApplicationDetailResponse{
		Response:    NewResponse("Application View", user, header),
		Application: app,
	}
}

// ApplicationFormResponse is the struct for the application form responses.
type ApplicationFormResponse struct {
	*ApplicationDetailResponse

	Clients      *model.Clients
	Projects     *model.Projects
	Environments []*model.Environment
	Databases    *model.Databases
	Runtimes     []*model.Runtime
	Pools        *model.Pools
	Domains      []*model.Domain
	CurrentUser  *model.User
}

// NewApplicationFormResponse is a constructor for the ApplicationFormResponse struct.
func NewApplicationFormResponse(
	title string,
	user *model.User,
	app *model.Application,
	clients *model.Clients,
	projects *model.Projects,
	envs []*model.Environment,
	dbs *model.Databases,
	runtimes []*model.Runtime,
	pools *model.Pools,
	domains []*model.Domain,
) *ApplicationFormResponse {
	appDetailResponse := NewApplicationDetailResponse(user, app)
	appDetailResponse.Title = title
	appDetailResponse.Header.Title = title
	appDetailResponse.Header.Buttons = []*ActionButton{}
	return &ApplicationFormResponse{
		ApplicationDetailResponse: appDetailResponse,
		Clients:                   clients,
		Projects:                  projects,
		Environments:              envs,
		Databases:                 dbs,
		Runtimes:                  runtimes,
		Pools:                     pools,
		Domains:                   domains,
		CurrentUser:               user,
	}
}

// ApplicationListResponse is the struct for the application list page.
type ApplicationListResponse struct {
	*Response
	Applications []*model.Application
}

// NewApplicationListResponse is a constructor for the ApplicationListResponse struct.
func NewApplicationListResponse(user *model.User, apps []*model.Application) *ApplicationListResponse {
	header := &HeaderBlock{
		Title:       "Application List",
		CurrentUser: user,
		Buttons: []*ActionButton{
			{
				Label:     "New",
				Link:      "/admin/application/create",
				Privilege: "applications.create",
			},
		},
	}
	return &ApplicationListResponse{
		Response:     NewResponse("Application List", user, header),
		Applications: apps,
	}
}

// ApplicationImportToEnvironmentFormResponse is the struct for the import application to environment form responses.
type ApplicationImportToEnvironmentFormResponse struct {
	*Response
	Environment *model.Environment
}

// NewApplicationImportToEnvironmentFormResponse is a constructor for the ApplicationImportToEnvironmentFormResponse struct.
func NewApplicationImportToEnvironmentFormResponse(user *model.User, env *model.Environment) *ApplicationImportToEnvironmentFormResponse {
	header := &HeaderBlock{
		Title:       "Import Application to Environment",
		CurrentUser: user,
		Buttons:     []*ActionButton{},
	}
	return &ApplicationImportToEnvironmentFormResponse{
		Response:    NewResponse("Import Application to Environment", user, header),
		Environment: env,
	}
}

// ApplicationMappingToEnvironmentFormResponse is the struct for the mapping application to environment form responses.
type ApplicationMappingToEnvironmentFormResponse struct {
	*Response
	Environment *model.Environment
	FileID      string
}

// NewApplicationMappingToEnvironmentFormResponse is a constructor for the ApplicationMappingToEnvironmentFormResponse struct.
func NewApplicationMappingToEnvironmentFormResponse(user *model.User, env *model.Environment, fileID string) *ApplicationMappingToEnvironmentFormResponse {
	header := &HeaderBlock{
		Title:       "Import Mapping to Environment",
		CurrentUser: user,
		Buttons:     []*ActionButton{},
	}
	return &ApplicationMappingToEnvironmentFormResponse{
		Response:    NewResponse("Import Mapping to Environment", user, header),
		Environment: env,
		FileID:      fileID,
	}
}

// ApplicationImportRowResult is the struct for the application import row results.
type ApplicationImportRowResult struct {
	ErrorMessage string
	RowData      []string
	Application  *model.Application
}

// ApplicationImportToEnvironmentListResponse is the struct for the import application to environment list page.
type ApplicationImportToEnvironmentListResponse struct {
	*Response
	Environment *model.Environment
	FileID      string
	Result      map[int]*ApplicationImportRowResult
}

// NewApplicationImportToEnvironmentListResponse is a constructor for the ApplicationImportToEnvironmentListResponse struct.
func NewApplicationImportToEnvironmentListResponse(user *model.User, env *model.Environment, fileID string, result map[int]*ApplicationImportRowResult) *ApplicationImportToEnvironmentListResponse {
	header := &HeaderBlock{
		Title:       "Import Application to Environment",
		CurrentUser: user,
		Buttons:     []*ActionButton{},
	}
	return &ApplicationImportToEnvironmentListResponse{
		Response:    NewResponse("Import Application to Environment", user, header),
		Environment: env,
		FileID:      fileID,
		Result:      result,
	}
}
