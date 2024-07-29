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
	FormItems []*FormItem
}

// NewApplicationFormResponse is a constructor for the ApplicationFormResponse struct.
func NewApplicationFormResponse(
	title string,
	user *model.User,
	app *model.Application,
	clients *model.Clients,
	projects *model.Projects,
	envs *model.Environments,
	dbs *model.Databases,
	runtimes *model.Runtimes,
	pools *model.Pools,
	domains *model.Domains,
) *ApplicationFormResponse {
	appDetailResponse := NewApplicationDetailResponse(user, app)
	appDetailResponse.Title = title
	appDetailResponse.Header.Title = title
	appDetailResponse.Header.Buttons = []*ActionButton{{Label: "Back", Link: "/admin/application/list", Privilege: "applications.view"}}
	selectedDomains := SelectedOptions{}
	if app.Domains != nil {
		for _, domain := range app.Domains {
			selectedDomains = append(selectedDomains, domain.ID)
		}
	}
	formItems := []*FormItem{
		// Client.
		{Label: "Client", Type: "select", Name: "client", Options: clients.ToMap(), Required: true},
		// Project.
		{Label: "Project", Type: "select", Name: "project", Options: projects.ToMap(), Required: true},
		// Environment.
		{Label: "Environment", Type: "select", Name: "environment", Options: envs.ToMap(), Required: true},
		// Database.
		{Label: "Database", Type: "select", Name: "database", Options: dbs.ToMap(), Required: true},
		// DB Name.
		{Label: "DB Name", Type: "text", Name: "db_name", Value: app.DBName, Required: true},
		// DB User.
		{Label: "DB User", Type: "text", Name: "db_user", Value: app.DBUser, Required: true},
		// Runtime.
		{Label: "Runtime", Type: "select", Name: "runtime", Options: runtimes.ToMap(), Required: true},
		// Pool.
		{Label: "Pool", Type: "select", Name: "pool", Options: pools.ToMap(), Required: true},
		// Repo URL.
		{Label: "Repository", Type: "text", Name: "repository", Value: app.Repository},
		// Branch.
		{Label: "Branch", Type: "text", Name: "branch", Value: app.Branch},
		// Framework
		{Label: "Framework", Type: "text", Name: "framework", Value: app.Framework},
		// Document root.
		{Label: "Document Root", Type: "text", Name: "document_root", Value: app.DocumentRoot},
		// Domains.
		{Label: "Domains", Type: "checkboxgroup", Name: "domains", Options: domains.ToMap(), SelectedOptions: selectedDomains},
	}
	return &ApplicationFormResponse{
		ApplicationDetailResponse: appDetailResponse,
		FormItems:                 formItems,
	}
}

// ApplicationListResponse is the struct for the application list page.
type ApplicationListResponse struct {
	*Response
	Applications *model.Applications
}

// NewApplicationListResponse is a constructor for the ApplicationListResponse struct.
func NewApplicationListResponse(user *model.User, apps *model.Applications) *ApplicationListResponse {
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
