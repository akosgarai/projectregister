package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewApplicationDetailResponse is a constructor for the DetailResponse struct for an application.
func NewApplicationDetailResponse(user *model.User, app *model.Application) *DetailResponse {
	headerText := "Application Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if user.HasPrivilege("applications.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/application/update/%d", app.ID)))
	}
	if user.HasPrivilege("applications.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/application/delete/%d", app.ID)))
	}
	if user.HasPrivilege("applications.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/application/list"))
	}
	dbValues := &components.DetailValues{
		{Value: app.Database.Name, Link: fmt.Sprintf("/admin/database/view/%d", app.Database.ID)},
	}
	if app.DBName != "" {
		*dbValues = append(*dbValues, &components.DetailValue{Value: app.DBUser + " / " + app.DBName})
	}
	codebaseValues := &components.DetailValues{}
	if app.Repository != "" {
		value := app.Repository
		if app.Branch != "" {
			value = fmt.Sprintf("%s (%s)", app.Repository, app.Branch)
		}
		*codebaseValues = append(*codebaseValues, &components.DetailValue{Value: value, Link: app.Repository})
	}
	domainValues := &components.DetailValues{}
	if app.Domains != nil {
		for _, domain := range app.Domains {
			*domainValues = append(*domainValues, &components.DetailValue{Value: domain.Name, Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)})
		}
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", app.ID)}}},
		{Label: "Client", Value: &components.DetailValues{{Value: app.Client.Name, Link: fmt.Sprintf("/admin/client/view/%d", app.Client.ID)}}},
		{Label: "Project", Value: &components.DetailValues{{Value: app.Project.Name, Link: fmt.Sprintf("/admin/project/view/%d", app.Project.ID)}}},
		{Label: "Environment", Value: &components.DetailValues{{Value: app.Environment.Name, Link: fmt.Sprintf("/admin/environment/view/%d", app.Environment.ID)}}},
		{Label: "Database", Value: dbValues},
		{Label: "Runtime", Value: &components.DetailValues{{Value: app.Runtime.Name, Link: fmt.Sprintf("/admin/runtime/view/%d", app.Runtime.ID)}}},
		{Label: "Pool", Value: &components.DetailValues{{Value: app.Pool.Name, Link: fmt.Sprintf("/admin/pool/view/%d", app.Pool.ID)}}},
		{Label: "Codebase", Value: codebaseValues},
		{Label: "Framework", Value: &components.DetailValues{{Value: app.Framework}}},
		{Label: "Document Root", Value: &components.DetailValues{{Value: app.DocumentRoot}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: app.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: app.UpdatedAt}}},
		{Label: "Domains", Value: domainValues},
	}
	return NewDetailResponse(headerText, user, headerContent, details)
}

// ApplicationFormResponse is the struct for the application form responses.
type ApplicationFormResponse struct {
	*DetailResponse
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
	appDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/application/list")}
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
		DetailResponse: appDetailResponse,
		FormItems:      formItems,
	}
}

// NewApplicationListResponse is a constructor for the ListingResponse struct of the applications.
func NewApplicationListResponse(user *model.User, apps *model.Applications) *ListingResponse {
	headerText := "Application List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if user.HasPrivilege("applications.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/application/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Client", "Project", "Environment", "Database", "Runtime", "Pool", "Codebase", "Framework", "Document Root", "Domains", "Created At", "Updated At", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := user.HasPrivilege("applications.update")
	userCanDelete := user.HasPrivilege("applications.delete")
	for _, application := range *apps {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", application.ID)}}}
		columns = append(columns, idColumn)
		clientViewLink := ""
		if user.HasPrivilege("clients.view") {
			clientViewLink = fmt.Sprintf("/admin/client/view/%d", application.Client.ID)
		}
		clientColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Client.Name, Link: clientViewLink}}}
		columns = append(columns, clientColumn)
		projectViewLink := ""
		if user.HasPrivilege("projects.view") {
			projectViewLink = fmt.Sprintf("/admin/project/view/%d", application.Project.ID)
		}
		projectColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Project.Name, Link: projectViewLink}}}
		columns = append(columns, projectColumn)
		environmentViewLink := ""
		if user.HasPrivilege("environments.view") {
			environmentViewLink = fmt.Sprintf("/admin/environment/view/%d", application.Environment.ID)
		}
		environmentColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Environment.Name, Link: environmentViewLink}}}
		columns = append(columns, environmentColumn)
		databaseViewLink := ""
		if user.HasPrivilege("databases.view") {
			databaseViewLink = fmt.Sprintf("/admin/database/view/%d", application.Database.ID)
		}
		databaseColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Database.Name, Link: databaseViewLink}}}
		if application.DBName != "" {
			*databaseColumn.Values = append(*databaseColumn.Values, &components.ListingColumnValue{Value: application.DBUser + " / " + application.DBName})
		}
		columns = append(columns, databaseColumn)
		runtimeViewLink := ""
		if user.HasPrivilege("runtimes.view") {
			runtimeViewLink = fmt.Sprintf("/admin/runtime/view/%d", application.Runtime.ID)
		}
		runtimeColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Runtime.Name, Link: runtimeViewLink}}}
		columns = append(columns, runtimeColumn)
		poolViewLink := ""
		if user.HasPrivilege("pools.view") {
			poolViewLink = fmt.Sprintf("/admin/pool/view/%d", application.Pool.ID)
		}
		poolColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Pool.Name, Link: poolViewLink}}}
		columns = append(columns, poolColumn)
		codebaseColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Repository, Link: application.Repository}}}
		if application.Branch != "" {
			*codebaseColumn.Values = append(*codebaseColumn.Values, &components.ListingColumnValue{Value: application.Branch})
		}
		columns = append(columns, codebaseColumn)
		frameworkColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Framework}}}
		columns = append(columns, frameworkColumn)
		documentRootColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.DocumentRoot}}}
		columns = append(columns, documentRootColumn)
		domainsColumn := &components.ListingColumn{Values: &components.ListingColumnValues{}}
		if application.Domains != nil {
			for _, domain := range application.Domains {
				*domainsColumn.Values = append(*domainsColumn.Values, &components.ListingColumnValue{Value: domain.Name, Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)})
			}
		}
		columns = append(columns, domainsColumn)
		createAtColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.CreatedAt}}}
		columns = append(columns, createAtColumn)
		updateAtColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.UpdatedAt}}}
		columns = append(columns, updateAtColumn)

		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/application/view/%d", application.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/application/update/%d", application.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/application/delete/%d", application.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, user, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}

// ApplicationImportToEnvironmentFormResponse is the struct for the import application to environment form responses.
type ApplicationImportToEnvironmentFormResponse struct {
	*Response
	Environment *model.Environment
}

// NewApplicationImportToEnvironmentFormResponse is a constructor for the ApplicationImportToEnvironmentFormResponse struct.
func NewApplicationImportToEnvironmentFormResponse(user *model.User, env *model.Environment) *ApplicationImportToEnvironmentFormResponse {
	headerText := "Import Application to Environment"
	header := components.NewContentHeader(headerText, []*components.Link{})
	return &ApplicationImportToEnvironmentFormResponse{
		Response:    NewResponse(headerText, user, header),
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
	headerText := "Import Mapping to Environment"
	header := components.NewContentHeader(headerText, []*components.Link{})
	return &ApplicationMappingToEnvironmentFormResponse{
		Response:    NewResponse(headerText, user, header),
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
	headerText := "Import Application to Environment"
	header := components.NewContentHeader(headerText, []*components.Link{})
	return &ApplicationImportToEnvironmentListResponse{
		Response:    NewResponse(headerText, user, header),
		Environment: env,
		FileID:      fileID,
		Result:      result,
	}
}
