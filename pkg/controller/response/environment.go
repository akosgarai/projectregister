package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// EnvironmentResponse is the struct for the environment page.
type EnvironmentResponse struct {
	*Response
	Environment *model.Environment
}

// EnvironmentDetailResponse is the struct for the environment detail page.
type EnvironmentDetailResponse struct {
	*EnvironmentResponse
	Details *DetailItems
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
			{
				Label:     "List",
				Link:      "/admin/environment/list",
				Privilege: "environments.view",
			},
		},
	}
	serverValues := DetailValues{}
	if environment.Servers != nil {
		for _, server := range environment.Servers {
			serverValues = append(serverValues, &DetailValue{Value: server.Name, Link: fmt.Sprintf("/admin/server/view/%d", server.ID)})
		}
	}
	dbValues := DetailValues{}
	if environment.Databases != nil {
		for _, db := range environment.Databases {
			dbValues = append(dbValues, &DetailValue{Value: db.Name, Link: fmt.Sprintf("/admin/database/view/%d", db.ID)})
		}
	}

	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", environment.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: environment.Name}}},
		{Label: "Description", Value: &DetailValues{{Value: environment.Description}}},
		{Label: "Created At", Value: &DetailValues{{Value: environment.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: environment.UpdatedAt}}},
		{Label: "Servers", Value: &serverValues},
		{Label: "Databases", Value: &dbValues},
	}
	return &EnvironmentDetailResponse{
		EnvironmentResponse: &EnvironmentResponse{
			Response:    NewResponse("Environment Detail", currentUser, header),
			Environment: environment,
		},
		Details: details,
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
		{Label: "Servers", Type: "checkboxgroup", Name: "servers", Options: servers.ToMap(), SelectedOptions: selectedServers},
		// Databases.
		{Label: "Databases", Type: "checkboxgroup", Name: "databases", Options: databases.ToMap(), SelectedOptions: selectedDatabases},
	}
	return &EnvironmentFormResponse{
		EnvironmentDetailResponse: environmentDetailResponse,
		FormItems:                 formItems,
	}
}

// EnvironmentListResponse is the struct for the environment list page.
type EnvironmentListResponse struct {
	*Response
	Listing *Listing
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
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Description", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("environments.update")
	userCanDelete := currentUser.HasPrivilege("environments.delete")
	for _, environment := range *environments {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", environment.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: environment.Name}}}
		columns = append(columns, nameColumn)
		desctiptionColumn := &ListingColumn{&ListingColumnValues{{Value: environment.Description}}}
		columns = append(columns, desctiptionColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/environment/detail/%d", environment.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/environment/update/%d", environment.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/environment/delete/%d", environment.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &EnvironmentListResponse{
		Response: NewResponse("Environment List", currentUser, header),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
