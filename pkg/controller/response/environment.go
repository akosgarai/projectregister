package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewEnvironmentDetailResponse is a constructor for the DetailResponse struct for an environment.
func NewEnvironmentDetailResponse(currentUser *model.User, environment *model.Environment) *DetailResponse {
	headerText := "Environment Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "environments", fmt.Sprintf("%d", environment.ID)))
	if currentUser.HasPrivilege("applications.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Import Applications", fmt.Sprintf("/admin/application/import-to-environment/%d", environment.ID)))
	}
	serverValues := components.DetailValues{}
	if environment.Servers != nil {
		for _, server := range environment.Servers {
			serverValues = append(serverValues, &components.DetailValue{Value: server.Name, Link: fmt.Sprintf("/admin/server/view/%d", server.ID)})
		}
	}
	dbValues := components.DetailValues{}
	if environment.Databases != nil {
		for _, db := range environment.Databases {
			dbValues = append(dbValues, &components.DetailValue{Value: db.Name, Link: fmt.Sprintf("/admin/database/view/%d", db.ID)})
		}
	}

	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", environment.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: environment.Name}}},
		{Label: "Description", Value: &components.DetailValues{{Value: environment.Description}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: environment.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: environment.UpdatedAt}}},
		{Label: "Servers", Value: &serverValues},
		{Label: "Databases", Value: &dbValues},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateEnvironmentResponse is a constructor for the FormResponse struct for creating a new environment.
func NewCreateEnvironmentResponse(currentUser *model.User, servers *model.Servers, databases *model.Databases) *FormResponse {
	return newEnvironmentFormResponse("Create Environment", currentUser, &model.Environment{}, servers, databases, "/admin/environment/create", "POST", "Create")
}

// NewUpdateEnvironmentResponse is a constructor for the FormResponse struct for updating an environment.
func NewUpdateEnvironmentResponse(currentUser *model.User, environment *model.Environment, servers *model.Servers, databases *model.Databases) *FormResponse {
	return newEnvironmentFormResponse("Update Environment", currentUser, environment, servers, databases, fmt.Sprintf("/admin/environment/update/%d", environment.ID), "POST", "Update")
}

// newEnvironmentFormResponse is a constructor for the FormResponse struct for an environment.
func newEnvironmentFormResponse(title string, currentUser *model.User, environment *model.Environment, servers *model.Servers, databases *model.Databases, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/environment/list")})
	var selectedServers, selectedDatabases []int64
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

	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", environment.Name, true, nil, nil),
		// Description.
		components.NewFormItem("Description", "description", "textarea", environment.Description, false, nil, nil),
		// Servers.
		components.NewFormItem("Servers", "servers", "checkboxgroup", "", false, servers.ToMap(), selectedServers),
		// Databases.
		components.NewFormItem("Databases", "databases", "checkboxgroup", "", false, databases.ToMap(), selectedDatabases),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewEnvironmentListResponse is a constructor for the ListingResponse struct of the environments.
func NewEnvironmentListResponse(currentUser *model.User, environments *model.Environments, servers *model.Servers, databases *model.Databases) *ListingResponse {
	headerText := "Environment List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("environments.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/environment/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Description", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("environments.update")
	userCanDelete := currentUser.HasPrivilege("environments.delete")
	for _, environment := range *environments {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", environment.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: environment.Name}}}
		columns = append(columns, nameColumn)
		desctiptionColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: environment.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/environment/view/%d", environment.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/environment/update/%d", environment.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/environment/delete/%d", environment.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	/* Create the search form. The only form item is the name. */
	formItems := []*components.FormItem{
		components.NewFormItem("Name", "name", "text", "", false, nil, nil),
		components.NewFormItem("Description", "description", "text", "", false, nil, nil),
		components.NewFormItem("Server", "server", "multiselect", "", false, servers.ToMap(), nil),
		components.NewFormItem("Database", "database", "multiselect", "", false, databases.ToMap(), nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/environment/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
