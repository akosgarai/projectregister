package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// DatabaseResponse is the struct for the database page.
type DatabaseResponse struct {
	*Response
	Database *model.Database
}

// DatabaseDetailResponse is the struct for the database detail page.
type DatabaseDetailResponse struct {
	*DatabaseResponse
	Details *DetailItems
}

// NewDatabaseDetailResponse is a constructor for the DatabaseDetailResponse struct.
func NewDatabaseDetailResponse(currentUser *model.User, database *model.Database) *DatabaseDetailResponse {
	header := &HeaderBlock{
		Title:       "Database Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/database/update/%d", database.ID),
				Privilege: "databases.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/database/delete/%d", database.ID),
				Privilege: "databases.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/database/list",
				Privilege: "databases.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", database.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: database.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: database.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: database.UpdatedAt}}},
	}
	return &DatabaseDetailResponse{
		DatabaseResponse: &DatabaseResponse{
			Response: NewResponse("Database Detail", currentUser, header),
			Database: database,
		},
		Details: details,
	}
}

// DatabaseFormResponse is the struct for the database form responses.
type DatabaseFormResponse struct {
	*DatabaseDetailResponse
	FormItems []*FormItem
}

// NewDatabaseFormResponse is a constructor for the DatabaseFormResponse struct.
func NewDatabaseFormResponse(title string, currentUser *model.User, database *model.Database) *DatabaseFormResponse {
	databaseDetailResponse := NewDatabaseDetailResponse(currentUser, database)
	databaseDetailResponse.Header.Title = title
	databaseDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	databaseDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/database/list"), Privilege: "databases.view"}}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: database.Name, Required: true},
	}
	return &DatabaseFormResponse{
		DatabaseDetailResponse: databaseDetailResponse,
		FormItems:              formItems,
	}
}

// DatabaseListResponse is the struct for the database list page.
type DatabaseListResponse struct {
	*Response
	Databases *model.Databases
}

// NewDatabaseListResponse is a constructor for the DatabaseListResponse struct.
func NewDatabaseListResponse(currentUser *model.User, databases *model.Databases) *DatabaseListResponse {
	header := &HeaderBlock{
		Title:       "Database List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/database/create",
				Privilege: "databases.create",
			},
		},
	}
	return &DatabaseListResponse{
		Response:  NewResponse("Database List", currentUser, header),
		Databases: databases,
	}
}
