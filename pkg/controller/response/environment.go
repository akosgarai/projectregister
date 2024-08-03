package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewEnvironmentDetailResponse is a constructor for the DetailResponse struct for an environment.
func NewEnvironmentDetailResponse(currentUser *model.User, environment *model.Environment) *DetailResponse {
	headerText := "Environment Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("environments.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/environment/update/%d", environment.ID)))
	}
	if currentUser.HasPrivilege("environments.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/environment/delete/%d", environment.ID)))
	}
	if currentUser.HasPrivilege("environments.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/environment/list"))
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

// EnvironmentFormResponse is the struct for the environment form responses.
type EnvironmentFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewEnvironmentFormResponse is a constructor for the EnvironmentFormResponse struct.
func NewEnvironmentFormResponse(title string, currentUser *model.User, environment *model.Environment, servers *model.Servers, databases *model.Databases) *EnvironmentFormResponse {
	environmentDetailResponse := NewEnvironmentDetailResponse(currentUser, environment)
	environmentDetailResponse.Header.Title = title
	environmentDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	environmentDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/environment/list")}
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
		{Label: "Servers", Type: "checkboxgroup", Name: "servers", Options: servers.ToMap(), SelectedOptions: selectedServers},
		// Databases.
		{Label: "Databases", Type: "checkboxgroup", Name: "databases", Options: databases.ToMap(), SelectedOptions: selectedDatabases},
	}
	return &EnvironmentFormResponse{
		DetailResponse: environmentDetailResponse,
		FormItems:      formItems,
	}
}

// NewEnvironmentListResponse is a constructor for the ListingResponse struct of the environments.
func NewEnvironmentListResponse(currentUser *model.User, environments *model.Environments) *ListingResponse {
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
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", environment.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: environment.Name}}}
		columns = append(columns, nameColumn)
		desctiptionColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: environment.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
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
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
