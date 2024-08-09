package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/parser"
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

// NewCreateApplicationResponse is a constructor for the FormResponse struct for the application create page.
func NewCreateApplicationResponse(
	currentUser *model.User,
	clients *model.Clients,
	projects *model.Projects,
	envs *model.Environments,
	dbs *model.Databases,
	runtimes *model.Runtimes,
	pools *model.Pools,
	domains *model.Domains,
) *FormResponse {
	return newApplicationFormResponse("Create Application", currentUser, &model.Application{}, clients, projects, envs, dbs, runtimes, pools, domains, "/admin/application/create", "POST", "Create")
}

// NewUpdateApplicationResponse is a constructor for the FormResponse struct for the application update page.
func NewUpdateApplicationResponse(
	currentUser *model.User,
	app *model.Application,
	clients *model.Clients,
	projects *model.Projects,
	envs *model.Environments,
	dbs *model.Databases,
	runtimes *model.Runtimes,
	pools *model.Pools,
	domains *model.Domains,
) *FormResponse {
	return newApplicationFormResponse("Update Application", currentUser, app, clients, projects, envs, dbs, runtimes, pools, domains, fmt.Sprintf("/admin/application/update/%d", app.ID), "POST", "Update")
}

// NewApplicationFormResponse is a constructor for the ApplicationFormResponse struct.
func newApplicationFormResponse(
	title string,
	currentUser *model.User,
	app *model.Application,
	clients *model.Clients,
	projects *model.Projects,
	envs *model.Environments,
	dbs *model.Databases,
	runtimes *model.Runtimes,
	pools *model.Pools,
	domains *model.Domains,
	action, method, submitLabel string,
) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/application/list")})

	var selectedClients, selectedProjects, selectedEnvironments, selectedDatabases, selectedRuntimes, selectedPools, selectedDomains []int64

	if app.Client != nil {
		selectedClients = []int64{app.Client.ID}
	}
	if app.Project != nil {
		selectedProjects = []int64{app.Project.ID}
	}
	if app.Environment != nil {
		selectedEnvironments = []int64{app.Environment.ID}
	}
	if app.Database != nil {
		selectedDatabases = []int64{app.Database.ID}
	}
	if app.Runtime != nil {
		selectedRuntimes = []int64{app.Runtime.ID}
	}
	if app.Pool != nil {
		selectedPools = []int64{app.Pool.ID}
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
	formItems := []*components.FormItem{
		// Client.
		components.NewFormItem("Client", "client", "select", "", true, clients.ToMap(), selectedClients),
		// Project.
		components.NewFormItem("Project", "project", "select", "", true, projects.ToMap(), selectedProjects),
		// Environment.
		components.NewFormItem("Environment", "environment", "select", "", true, envs.ToMap(), selectedEnvironments),
		// Database.
		components.NewFormItem("Database", "database", "select", "", true, dbs.ToMap(), selectedDatabases),
		// DB Name.
		components.NewFormItem("DB Name", "db_name", "text", app.DBName, true, nil, nil),
		// DB User.
		components.NewFormItem("DB User", "db_user", "text", app.DBUser, true, nil, nil),
		// Runtime.
		components.NewFormItem("Runtime", "runtime", "select", "", true, runtimes.ToMap(), selectedRuntimes),
		// Pool.
		components.NewFormItem("Pool", "pool", "select", "", true, pools.ToMap(), selectedPools),
		// Repo URL.
		components.NewFormItem("Repository", "repository", "text", app.Repository, false, nil, nil),
		// Branch.
		components.NewFormItem("Branch", "branch", "text", app.Branch, false, nil, nil),
		// Framework
		components.NewFormItem("Framework", "framework", "text", app.Framework, false, nil, nil),
		// Document root.
		components.NewFormItem("Document Root", "document_root", "text", app.DocumentRoot, false, nil, nil),
		// Domains.
		components.NewFormItem("Domains", "domains", "checkboxgroup", "", false, domains.ToMap(), selectedDomains),
	}
	form := &components.Form{
		Method: method,
		Action: action,
		Items:  formItems,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
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

// NewApplicationImportToEnvironmentFormResponse is a constructor for the ApplicationImportToEnvironmentFormResponse struct.
func NewApplicationImportToEnvironmentFormResponse(currentUser *model.User, env *model.Environment) *FormResponse {
	headerText := "Import Application to Environment"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	checkboxOptions := map[int64]string{
		1: "Yes",
	}
	formItems := []*components.FormItem{
		components.NewFormItem("CSV File", "csvfile", "file", "", true, nil, nil),
		components.NewFormItem("Has Header", "has_header", "checkboxgroup", "true", false, checkboxOptions, nil),
	}
	form := &components.Form{
		Items:     formItems,
		Action:    fmt.Sprintf("/admin/application/import-to-environment/%d", env.ID),
		Method:    "POST",
		Submit:    "Upload",
		Multipart: true,
	}
	return NewFormResponse(headerText, currentUser, headerContent, form)
}

// NewApplicationMappingToEnvironmentFormResponse is a constructor for the ApplicationMappingToEnvironmentFormResponse struct.
func NewApplicationMappingToEnvironmentFormResponse(currentUser *model.User, env *model.Environment, fileID string, data [][]string, headers []string) *ApplicationImportMappingResponse {
	headerText := "Import Mapping to Environment"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	// On case of the headers are not set, we need to set them.
	// Indexes are starting from 1.
	if len(headers) == 0 {
		for i := range data[0] {
			headers = append(headers, fmt.Sprintf("Column %d", i+1))
		}
	}
	rows := components.ListingRows{}
	// Add 5 rows to the listing as preview.
	for i := 0; i < 5; i++ {
		columns := components.ListingColumns{}
		for _, dataItem := range data[i] {
			columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: dataItem}}})
		}
		rows = append(rows, &components.ListingRow{Columns: &columns})
	}

	listing := &components.Listing{
		Header: &components.ListingHeader{Headers: headers},
		Rows:   &rows,
	}
	formMapping := parser.NewApplicationImportMapping()
	formItems := []*components.FormItem{
		components.NewFormItem("", "environment_id", "hidden", fmt.Sprintf("%d", env.ID), true, nil, nil),
		components.NewFormItem("", "file_id", "hidden", fileID, true, nil, nil),
	}
	// Select the headers.
	mappingOptions := map[int64]string{}
	for i, header := range headers {
		mappingOptions[int64(i)] = header
	}
	// Add the mapping rules to the form items.
	formItemsInOrder := []string{"client", "project", "runtime", "pool", "domains", "framework", "database", "database_name", "database_user", "doc_root", "repository", "branch"}
	for _, header := range formItemsInOrder {
		// on case of we have header in the headers (case insensitive match), we set the column index.
		// Then add 2 form items. One for the column index (select, options are the headers), and one for the custom value (text input).
		// On case of the header is not in the headers, we keep the original value.
		for i, headerItem := range headers {
			if headerItem == header {
				formMapping[header].ColumnIndex = i
				break
			}
		}
		if formMapping[header].ColumnIndex == -1 {
			formItems = append(formItems, components.NewFormItem(header, header, "select", "", false, mappingOptions, nil))
		} else {
			formItems = append(formItems, components.NewFormItem(header, header, "select", "", false, mappingOptions, []int64{int64(formMapping[header].ColumnIndex)}))
		}
		formItems = append(formItems, components.NewFormItem(fmt.Sprintf("Custom %s name", header), fmt.Sprintf("%s_custom", header), "text", "", false, nil, nil))
	}
	form := &components.Form{
		Items:     formItems,
		Action:    fmt.Sprintf("/admin/application/mapping-to-environment/%d/%s", env.ID, fileID),
		Method:    "POST",
		Submit:    "Upload",
		Multipart: true,
	}
	return NewApplicationImportMappingResponse(headerText, currentUser, headerContent, listing, form)
}

// NewApplicationImportToEnvironmentListResponse is a constructor for the ApplicationImportToEnvironmentListResponse struct.
func NewApplicationImportToEnvironmentListResponse(currentUser *model.User, env *model.Environment, fileID string, result *parser.ApplicationImportResult) *ListingResponse {
	headerText := "Import Application to Environment"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	listingHeader := &components.ListingHeader{
		Headers: []string{"Client", "Project", "Database", "Runtime", "Pool", "Codebase", "Framework", "Document Root", "Domains", "Status"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	for _, row := range *result {
		columns := components.ListingColumns{}
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["client"]}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["project"]}}})
		databaseText := row.RowData["database"]
		if row.RowData["db_name"] != "" {
			databaseText = fmt.Sprintf("%s %s / %s", row.RowData["database"], row.RowData["db_user"], row.RowData["db_name"])
		}
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: databaseText}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["runtime"]}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["pool"]}}})
		codebaseText := row.RowData["repository"]
		if row.RowData["branch"] != "" {
			codebaseText = fmt.Sprintf("%s (%s)", row.RowData["repository"], row.RowData["branch"])
		}
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: codebaseText}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["framework"]}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["doc_root"]}}})
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: row.RowData["domains"]}}})
		statusText := ""
		if row.Application != nil {
			statusText = "OK"
		} else {
			statusText = row.ErrorMessage
		}
		columns = append(columns, &components.ListingColumn{Values: &components.ListingColumnValues{{Value: statusText}}})

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
