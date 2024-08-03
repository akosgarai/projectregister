package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewDatabaseDetailResponse is a constructor for the DetailResponse struct for a database.
func NewDatabaseDetailResponse(currentUser *model.User, database *model.Database) *DetailResponse {
	headerText := "Database Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("databases.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/database/update/%d", database.ID)))
	}
	if currentUser.HasPrivilege("databases.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/database/delete/%d", database.ID)))
	}
	if currentUser.HasPrivilege("databases.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/database/list"))
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", database.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: database.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: database.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: database.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// DatabaseFormResponse is the struct for the database form responses.
type DatabaseFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewDatabaseFormResponse is a constructor for the DatabaseFormResponse struct.
func NewDatabaseFormResponse(title string, currentUser *model.User, database *model.Database) *DatabaseFormResponse {
	databaseDetailResponse := NewDatabaseDetailResponse(currentUser, database)
	databaseDetailResponse.Header.Title = title
	databaseDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	databaseDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/database/list")}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: database.Name, Required: true},
	}
	return &DatabaseFormResponse{
		DetailResponse: databaseDetailResponse,
		FormItems:      formItems,
	}
}

// NewDatabaseListResponse is a constructor for the ListingResponse struct of the databases.
func NewDatabaseListResponse(currentUser *model.User, databases *model.Databases) *ListingResponse {
	headerText := "Database List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("databases.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/database/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("databases.update")
	userCanDelete := currentUser.HasPrivilege("databases.delete")
	for _, database := range *databases {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", database.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: database.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/database/view/%d", database.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/database/update/%d", database.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/database/delete/%d", database.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
