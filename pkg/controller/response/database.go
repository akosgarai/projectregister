package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewDatabaseDetailResponse is a constructor for the DetailResponse struct for a database.
func NewDatabaseDetailResponse(currentUser *model.User, database *model.Database) *DetailResponse {
	headerText := "Database Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "databases", fmt.Sprintf("%d", database.ID)))
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", database.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: database.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: database.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: database.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateDatabaseResponse is a constructor for the FormResponse struct for creating a database.
func NewCreateDatabaseResponse(currentUser *model.User) *FormResponse {
	return newDatabaseFormResponse("Create Database", currentUser, &model.Database{}, "/admin/database/create", "POST", "Create")
}

// NewUpdateDatabaseResponse is a constructor for the FormResponse struct for updating a database.
func NewUpdateDatabaseResponse(currentUser *model.User, database *model.Database) *FormResponse {
	return newDatabaseFormResponse("Update Database", currentUser, database, fmt.Sprintf("/admin/database/update/%d", database.ID), "POST", "Update")
}

// newDatabaseFormResponse is a constructor for the FormResponse struct for a database.
func newDatabaseFormResponse(title string, currentUser *model.User, database *model.Database, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/database/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", database.Name, true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
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
