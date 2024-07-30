package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationResponse is the struct for the application page.
type ApplicationResponse struct {
	*Response
	Application *model.Application
}

// ApplicationDetailResponse is the struct for the application view response.
type ApplicationDetailResponse struct {
	*ApplicationResponse
	Details *DetailItems
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
			{
				Label:     "List",
				Link:      "/admin/application/list",
				Privilege: "applications.view",
			},
		},
	}
	dbValues := &DetailValues{
		{Value: app.Database.Name, Link: fmt.Sprintf("/admin/database/view/%d", app.Database.ID)},
	}
	if app.DBName != "" {
		*dbValues = append(*dbValues, &DetailValue{Value: app.DBUser + " / " + app.DBName})
	}
	codebaseValues := &DetailValues{}
	if app.Repository != "" {
		value := app.Repository
		if app.Branch != "" {
			value = fmt.Sprintf("%s (%s)", app.Repository, app.Branch)
		}
		*codebaseValues = append(*codebaseValues, &DetailValue{Value: value, Link: app.Repository})
	}
	domainValues := &DetailValues{}
	if app.Domains != nil {
		for _, domain := range app.Domains {
			*domainValues = append(*domainValues, &DetailValue{Value: domain.Name, Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)})
		}
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", app.ID)}}},
		{Label: "Client", Value: &DetailValues{{Value: app.Client.Name, Link: fmt.Sprintf("/admin/client/view/%d", app.Client.ID)}}},
		{Label: "Project", Value: &DetailValues{{Value: app.Project.Name, Link: fmt.Sprintf("/admin/project/view/%d", app.Project.ID)}}},
		{Label: "Environment", Value: &DetailValues{{Value: app.Environment.Name, Link: fmt.Sprintf("/admin/environment/view/%d", app.Environment.ID)}}},
		{Label: "Database", Value: dbValues},
		{Label: "Runtime", Value: &DetailValues{{Value: app.Runtime.Name, Link: fmt.Sprintf("/admin/runtime/view/%d", app.Runtime.ID)}}},
		{Label: "Pool", Value: &DetailValues{{Value: app.Pool.Name, Link: fmt.Sprintf("/admin/pool/view/%d", app.Pool.ID)}}},
		{Label: "Codebase", Value: codebaseValues},
		{Label: "Framework", Value: &DetailValues{{Value: app.Framework}}},
		{Label: "Document Root", Value: &DetailValues{{Value: app.DocumentRoot}}},
		{Label: "Created At", Value: &DetailValues{{Value: app.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: app.UpdatedAt}}},
		{Label: "Domains", Value: domainValues},
	}
	return &ApplicationDetailResponse{
		ApplicationResponse: &ApplicationResponse{
			Response:    NewResponse("Application View", user, header),
			Application: app,
		},
		Details: details,
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
	selectedClients := SelectedOptions{}
	selectedProjects := SelectedOptions{}
	selectedEnvironments := SelectedOptions{}
	selectedDatabases := SelectedOptions{}
	selectedRuntimes := SelectedOptions{}
	selectedPools := SelectedOptions{}
	selectedDomains := SelectedOptions{}
	if app.Client != nil {
		selectedClients = SelectedOptions{app.Client.ID}
	}
	if app.Project != nil {
		selectedProjects = SelectedOptions{app.Project.ID}
	}
	if app.Environment != nil {
		selectedEnvironments = SelectedOptions{app.Environment.ID}
	}
	if app.Database != nil {
		selectedDatabases = SelectedOptions{app.Database.ID}
	}
	if app.Runtime != nil {
		selectedRuntimes = SelectedOptions{app.Runtime.ID}
	}
	if app.Pool != nil {
		selectedPools = SelectedOptions{app.Pool.ID}
	}
	if app.Domains != nil {
		for _, domain := range app.Domains {
			selectedDomains = append(selectedDomains, domain.ID)
		}
	}
	// add the application domains to the domains list.
	for _, domain := range app.Domains {
		*domains = append(*domains, domain)
	}
	formItems := []*FormItem{
		// Client.
		{Label: "Client", Type: "select", Name: "client", Options: clients.ToMap(), SelectedOptions: selectedClients, Required: true},
		// Project.
		{Label: "Project", Type: "select", Name: "project", Options: projects.ToMap(), SelectedOptions: selectedProjects, Required: true},
		// Environment.
		{Label: "Environment", Type: "select", Name: "environment", Options: envs.ToMap(), SelectedOptions: selectedEnvironments, Required: true},
		// Database.
		{Label: "Database", Type: "select", Name: "database", Options: dbs.ToMap(), SelectedOptions: selectedDatabases, Required: true},
		// DB Name.
		{Label: "DB Name", Type: "text", Name: "db_name", Value: app.DBName, Required: true},
		// DB User.
		{Label: "DB User", Type: "text", Name: "db_user", Value: app.DBUser, Required: true},
		// Runtime.
		{Label: "Runtime", Type: "select", Name: "runtime", Options: runtimes.ToMap(), SelectedOptions: selectedRuntimes, Required: true},
		// Pool.
		{Label: "Pool", Type: "select", Name: "pool", Options: pools.ToMap(), SelectedOptions: selectedPools, Required: true},
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
	Listing *Listing
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
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Client", "Project", "Environment", "Database", "Runtime", "Pool", "Codebase", "Framework", "Document Root", "Domains", "Created At", "Updated At", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := user.HasPrivilege("applications.update")
	userCanDelete := user.HasPrivilege("applications.delete")
	for _, application := range *apps {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", application.ID)}}}
		columns = append(columns, idColumn)
		clientViewLink := ""
		if user.HasPrivilege("clients.view") {
			clientViewLink = fmt.Sprintf("/admin/client/view/%d", application.Client.ID)
		}
		clientColumn := &ListingColumn{&ListingColumnValues{{Value: application.Client.Name, Link: clientViewLink}}}
		columns = append(columns, clientColumn)
		projectViewLink := ""
		if user.HasPrivilege("projects.view") {
			projectViewLink = fmt.Sprintf("/admin/project/view/%d", application.Project.ID)
		}
		projectColumn := &ListingColumn{&ListingColumnValues{{Value: application.Project.Name, Link: projectViewLink}}}
		columns = append(columns, projectColumn)
		environmentViewLink := ""
		if user.HasPrivilege("environments.view") {
			environmentViewLink = fmt.Sprintf("/admin/environment/view/%d", application.Environment.ID)
		}
		environmentColumn := &ListingColumn{&ListingColumnValues{{Value: application.Environment.Name, Link: environmentViewLink}}}
		columns = append(columns, environmentColumn)
		databaseViewLink := ""
		if user.HasPrivilege("databases.view") {
			databaseViewLink = fmt.Sprintf("/admin/database/view/%d", application.Database.ID)
		}
		databaseColumn := &ListingColumn{&ListingColumnValues{{Value: application.Database.Name, Link: databaseViewLink}}}
		if application.DBName != "" {
			*databaseColumn.Values = append(*databaseColumn.Values, &ListingColumnValue{Value: application.DBUser + " / " + application.DBName})
		}
		columns = append(columns, databaseColumn)
		runtimeViewLink := ""
		if user.HasPrivilege("runtimes.view") {
			runtimeViewLink = fmt.Sprintf("/admin/runtime/view/%d", application.Runtime.ID)
		}
		runtimeColumn := &ListingColumn{&ListingColumnValues{{Value: application.Runtime.Name, Link: runtimeViewLink}}}
		columns = append(columns, runtimeColumn)
		poolViewLink := ""
		if user.HasPrivilege("pools.view") {
			poolViewLink = fmt.Sprintf("/admin/pool/view/%d", application.Pool.ID)
		}
		poolColumn := &ListingColumn{&ListingColumnValues{{Value: application.Pool.Name, Link: poolViewLink}}}
		columns = append(columns, poolColumn)
		codebaseColumn := &ListingColumn{&ListingColumnValues{{Value: application.Repository, Link: application.Repository}}}
		if application.Branch != "" {
			*codebaseColumn.Values = append(*codebaseColumn.Values, &ListingColumnValue{Value: application.Branch})
		}
		columns = append(columns, codebaseColumn)
		frameworkColumn := &ListingColumn{&ListingColumnValues{{Value: application.Framework}}}
		columns = append(columns, frameworkColumn)
		documentRootColumn := &ListingColumn{&ListingColumnValues{{Value: application.DocumentRoot}}}
		columns = append(columns, documentRootColumn)
		domainsColumn := &ListingColumn{&ListingColumnValues{}}
		if application.Domains != nil {
			for _, domain := range application.Domains {
				*domainsColumn.Values = append(*domainsColumn.Values, &ListingColumnValue{Value: domain.Name, Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)})
			}
		}
		columns = append(columns, domainsColumn)
		createAtColumn := &ListingColumn{&ListingColumnValues{{Value: application.CreatedAt}}}
		columns = append(columns, createAtColumn)
		updateAtColumn := &ListingColumn{&ListingColumnValues{{Value: application.UpdatedAt}}}
		columns = append(columns, updateAtColumn)

		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/application/detail/%d", application.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/application/update/%d", application.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/application/delete/%d", application.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &ApplicationListResponse{
		Response: NewResponse("Application List", user, header),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
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
