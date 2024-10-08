package response

import (
	"fmt"
	"math"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/parser"
	"github.com/akosgarai/projectregister/pkg/transformers"
)

// NewApplicationDetailResponse is a constructor for the DetailResponse struct for an application.
func NewApplicationDetailResponse(currentUser *model.User, app *model.Application) *DetailResponse {
	headerText := "Application Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "applications", fmt.Sprintf("%d", app.ID)))
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
		{Label: "Framework", Value: &components.DetailValues{{Value: app.Framework.Name, Link: fmt.Sprintf("/admin/framework/view/%d", app.Framework.ID)}}},
		{Label: "Document Root", Value: &components.DetailValues{{Value: app.DocumentRoot}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: app.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: app.UpdatedAt}}},
		{Label: "Domains", Value: domainValues},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
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
	frameworks *model.Frameworks,
	domains *model.Domains,
) *FormResponse {
	return newApplicationFormResponse("Create Application", currentUser, &model.Application{}, clients, projects, envs, dbs, runtimes, pools, frameworks, domains, "/admin/application/create", "POST", "Create")
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
	frameworks *model.Frameworks,
	domains *model.Domains,
) *FormResponse {
	return newApplicationFormResponse("Update Application", currentUser, app, clients, projects, envs, dbs, runtimes, pools, frameworks, domains, fmt.Sprintf("/admin/application/update/%d", app.ID), "POST", "Update")
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
	frameworks *model.Frameworks,
	domains *model.Domains,
	action, method, submitLabel string,
) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/application/list")})

	var selectedClients, selectedProjects, selectedEnvironments, selectedDatabases, selectedRuntimes, selectedPools, selectedFrameworks, selectedDomains []int64

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
	if app.Framework != nil {
		selectedFrameworks = []int64{app.Framework.ID}
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
		components.NewFormItem("Framework", "framework", "select", "", false, frameworks.ToMap(), selectedFrameworks),
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
func NewApplicationListResponse(
	currentUser *model.User,
	apps *model.Applications,
	clients *model.Clients,
	projects *model.Projects,
	envs *model.Environments,
	dbs *model.Databases,
	runtimes *model.Runtimes,
	pools *model.Pools,
	frameworks *model.Frameworks,
	filter *model.ApplicationFilter,
) *ListingResponse {
	headerText := "Application List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("applications.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/application/create"))
	}
	filteredHeaders := []string{}
	allColumns := filter.GetAllColumns()
	for _, header := range filter.VisibleColumns {
		filteredHeaders = append(filteredHeaders, allColumns[header])
	}
	listingHeader := &components.ListingHeader{
		Headers: filteredHeaders,
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("applications.update")
	userCanDelete := currentUser.HasPrivilege("applications.delete")
	for _, application := range *apps {
		columns := components.ListingColumns{}

		if filter.IsVisibleColumn("ID") {
			idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", application.ID)}}}
			columns = append(columns, idColumn)
		}
		if filter.IsVisibleColumn("Client") {
			clientViewLink := ""
			if currentUser.HasPrivilege("clients.view") {
				clientViewLink = fmt.Sprintf("/admin/client/view/%d", application.Client.ID)
			}
			clientColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Client.Name, Link: clientViewLink}}}
			columns = append(columns, clientColumn)
		}
		if filter.IsVisibleColumn("Project") {
			projectViewLink := ""
			if currentUser.HasPrivilege("projects.view") {
				projectViewLink = fmt.Sprintf("/admin/project/view/%d", application.Project.ID)
			}
			projectColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Project.Name, Link: projectViewLink}}}
			columns = append(columns, projectColumn)
		}
		if filter.IsVisibleColumn("Environment") {
			environmentViewLink := ""
			if currentUser.HasPrivilege("environments.view") {
				environmentViewLink = fmt.Sprintf("/admin/environment/view/%d", application.Environment.ID)
			}
			environmentColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Environment.Name, Link: environmentViewLink}}}
			columns = append(columns, environmentColumn)
		}
		if filter.IsVisibleColumn("Database") {
			databaseViewLink := ""
			if currentUser.HasPrivilege("databases.view") {
				databaseViewLink = fmt.Sprintf("/admin/database/view/%d", application.Database.ID)
			}
			databaseColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Database.Name, Link: databaseViewLink}}}
			if application.DBName != "" {
				*databaseColumn.Values = append(*databaseColumn.Values, &components.ListingColumnValue{Value: application.DBUser + " / " + application.DBName})
			}
			columns = append(columns, databaseColumn)
		}
		if filter.IsVisibleColumn("Runtime") {
			runtimeViewLink := ""
			if currentUser.HasPrivilege("runtimes.view") {
				runtimeViewLink = fmt.Sprintf("/admin/runtime/view/%d", application.Runtime.ID)
			}
			runtimeColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Runtime.Name, Link: runtimeViewLink}}}
			columns = append(columns, runtimeColumn)
		}
		if filter.IsVisibleColumn("Pool") {
			poolViewLink := ""
			if currentUser.HasPrivilege("pools.view") {
				poolViewLink = fmt.Sprintf("/admin/pool/view/%d", application.Pool.ID)
			}
			poolColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Pool.Name, Link: poolViewLink}}}
			columns = append(columns, poolColumn)
		}
		if filter.IsVisibleColumn("Codebase") {
			codebaseColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Repository, Link: application.Repository}}}
			if application.Branch != "" {
				*codebaseColumn.Values = append(*codebaseColumn.Values, &components.ListingColumnValue{Value: application.Branch})
			}
			columns = append(columns, codebaseColumn)
		}
		if filter.IsVisibleColumn("Framework") {
			frameViewLink := ""
			if currentUser.HasPrivilege("frameworks.view") {
				frameViewLink = fmt.Sprintf("/admin/framework/view/%d", application.Framework.ID)
			}
			frameworkColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.Framework.Name, Link: frameViewLink}}}
			columns = append(columns, frameworkColumn)
		}
		if filter.IsVisibleColumn("Document Root") {
			documentRootColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.DocumentRoot}}}
			columns = append(columns, documentRootColumn)
		}
		if filter.IsVisibleColumn("Score") {
			scoreValue := math.Min(float64(application.Framework.Score), math.Min(float64(application.Runtime.Score), float64(application.Environment.Score)))
			scoreColumn := &components.ListingColumn{Values: &components.ListingColumnValues{
				{Value: fmt.Sprintf("Environment Score: %d", application.Environment.Score)},
				{Value: fmt.Sprintf("Framework Score: %d", application.Framework.Score)},
				{Value: fmt.Sprintf("Runtime Score: %d", application.Runtime.Score)},
				{Value: fmt.Sprintf("Min Score: %d", int(scoreValue))}},
			}
			columns = append(columns, scoreColumn)
		}
		if filter.IsVisibleColumn("Domains") {
			domainsColumn := &components.ListingColumn{Values: &components.ListingColumnValues{}}
			if application.Domains != nil {
				for _, domain := range application.Domains {
					*domainsColumn.Values = append(*domainsColumn.Values, &components.ListingColumnValue{Value: domain.Name, Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)})
				}
			}
			columns = append(columns, domainsColumn)
		}
		if filter.IsVisibleColumn("Created At") {
			createAtColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.CreatedAt}}}
			columns = append(columns, createAtColumn)
		}
		if filter.IsVisibleColumn("Updated At") {
			updateAtColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: application.UpdatedAt}}}
			columns = append(columns, updateAtColumn)
		}

		if filter.IsVisibleColumn("Actions") {
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
		}

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	/* Create the search form. */
	formItems := []*components.FormItem{
		components.NewFormItem("Columns", "visible_columns", "multiselect", "", false, filter.GetAllColumns(), filter.VisibleColumns),
		components.NewFormItem("Client", "client", "multiselect", "", false, clients.ToMap(), transformers.StringSliceToInt64Slice(filter.ClientIDs)),
		components.NewFormItem("Project", "project", "multiselect", "", false, projects.ToMap(), transformers.StringSliceToInt64Slice(filter.ProjectIDs)),
		components.NewFormItem("Environment", "environment", "multiselect", "", false, envs.ToMap(), transformers.StringSliceToInt64Slice(filter.EnvironmentIDs)),
		components.NewFormItem("Database", "database", "multiselect", "", false, dbs.ToMap(), transformers.StringSliceToInt64Slice(filter.DatabaseIDs)),
		components.NewFormItem("DB User", "db_user", "text", filter.DBUser, false, nil, nil),
		components.NewFormItem("DB Name", "db_name", "text", filter.DBName, false, nil, nil),
		components.NewFormItem("Runtime", "runtime", "multiselect", "", false, runtimes.ToMap(), transformers.StringSliceToInt64Slice(filter.RuntimeIDs)),
		components.NewFormItem("Pool", "pool", "multiselect", "", false, pools.ToMap(), transformers.StringSliceToInt64Slice(filter.PoolIDs)),
		components.NewFormItem("Domain", "domain", "text", filter.Domain, false, nil, nil),
		components.NewFormItem("Repository", "repository", "text", filter.Repository, false, nil, nil),
		components.NewFormItem("Branch", "branch", "text", filter.Branch, false, nil, nil),
		components.NewFormItem("Framework", "framework", "multiselect", "", false, frameworks.ToMap(), transformers.StringSliceToInt64Slice(filter.FrameworkIDs)),
		components.NewFormItem("Document Root", "doc_root", "text", filter.DocRoot, false, nil, nil),
		components.NewFormItem("", "export-search", "submit", "Export", false, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/application/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
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
		// if the data index is missing, we are done, no need to continue.
		if len(data) <= i {
			break
		}
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
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, nil)
}
