package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewUserDetailResponse is a constructor for the DetailResponse struct for a user.
func NewUserDetailResponse(currentUser, user *model.User) *DetailResponse {
	headerText := "User Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("users.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/user/update/%d", user.ID)))
	}
	if currentUser.HasPrivilege("users.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/user/delete/%d", user.ID)))
	}
	if currentUser.HasPrivilege("users.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/user/list"))
	}
	roleLink := ""
	if currentUser.HasPrivilege("roles.view") {
		roleLink = fmt.Sprintf("/admin/role/view/%d", user.Role.ID)
	}
	roleValue := components.DetailValues{{Value: user.Role.Name, Link: roleLink}}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", user.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: user.Name}}},
		{Label: "Email", Value: &components.DetailValues{{Value: user.Email}}},
		{Label: "Role", Value: &roleValue},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// UserFormResponse is the struct for the user form responses.
type UserFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewUserFormResponse is a constructor for the UserFormResponse struct.
func NewUserFormResponse(title string, currentUser, user *model.User, roles *model.Roles) *UserFormResponse {
	userDetailResponse := NewUserDetailResponse(currentUser, user)
	userDetailResponse.Header.Title = title
	userDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	userDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/user/list")}
	roleID := ""
	selectedOptions := SelectedOptions{0}
	if user.Role != nil && user.Role.ID > 0 {
		roleID = fmt.Sprintf("%d", user.Role.ID)
		selectedOptions[0] = user.Role.ID
	}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: user.Name, Required: true},
		// Email.
		{Label: "Email", Type: "email", Name: "email", Value: user.Email, Required: true},
		// Password.
		{Label: "Password", Type: "password", Name: "password", Value: "", Required: false},
		// Roles.
		{
			Label:           "Role",
			Name:            "role",
			Type:            "select",
			Value:           roleID,
			Required:        true,
			Options:         roles.ToMap(),
			SelectedOptions: selectedOptions,
		},
	}
	return &UserFormResponse{
		DetailResponse: userDetailResponse,
		FormItems:      formItems,
	}
}

// NewUserListResponse is a constructor for the ListingResponse struct of the users.
func NewUserListResponse(currentUser *model.User, users []*model.User) *ListingResponse {
	headerText := "User List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("users.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/user/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Email", "Role", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanViewRoles := currentUser.HasPrivilege("roles.view")
	userCanUpdateUsers := currentUser.HasPrivilege("users.update")
	userCanDeleteUsers := currentUser.HasPrivilege("users.delete")
	for _, user := range users {
		columns := components.ListingColumns{}
		// ID
		idColumn := components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", user.ID)}}}
		columns = append(columns, &idColumn)
		// Name
		nameColumn := components.ListingColumn{&components.ListingColumnValues{{Value: user.Name}}}
		columns = append(columns, &nameColumn)
		// Email
		emailColumn := components.ListingColumn{&components.ListingColumnValues{{Value: user.Email}}}
		columns = append(columns, &emailColumn)
		// Role
		roleLink := ""
		if userCanViewRoles {
			roleLink = fmt.Sprintf("/admin/role/view/%d", user.Role.ID)
		}
		roleColumn := components.ListingColumn{&components.ListingColumnValues{{Value: user.Role.Name, Link: roleLink}}}
		// Actions
		columns = append(columns, &roleColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			// view link
			{Value: "View", Link: fmt.Sprintf("/admin/user/view/%d", user.ID)},
		}}
		if userCanUpdateUsers {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/user/update/%d", user.ID)})
		}
		if userCanDeleteUsers {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/user/delete/%d", user.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		row := components.ListingRow{Columns: &columns}
		listingRows = append(listingRows, &row)
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
