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

// UserListResponse is the struct for the user list page.
type UserListResponse struct {
	*Response
	Listing *Listing
}

// NewUserListResponse is a constructor for the UserListResponse struct.
func NewUserListResponse(currentUser *model.User, users []*model.User) *UserListResponse {
	headerText := "User List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("users.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/user/create"))
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Email", "Role", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanViewRoles := currentUser.HasPrivilege("roles.view")
	userCanUpdateUsers := currentUser.HasPrivilege("users.update")
	userCanDeleteUsers := currentUser.HasPrivilege("users.delete")
	for _, user := range users {
		columns := ListingColumns{}
		// ID
		idColumn := ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", user.ID)}}}
		columns = append(columns, &idColumn)
		// Name
		nameColumn := ListingColumn{&ListingColumnValues{{Value: user.Name}}}
		columns = append(columns, &nameColumn)
		// Email
		emailColumn := ListingColumn{&ListingColumnValues{{Value: user.Email}}}
		columns = append(columns, &emailColumn)
		// Role
		roleLink := ""
		if userCanViewRoles {
			roleLink = fmt.Sprintf("/admin/role/view/%d", user.Role.ID)
		}
		roleColumn := ListingColumn{&ListingColumnValues{{Value: user.Role.Name, Link: roleLink}}}
		// Actions
		columns = append(columns, &roleColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			// view link
			{Value: "View", Link: fmt.Sprintf("/admin/user/view/%d", user.ID)},
		}}
		if userCanUpdateUsers {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/user/update/%d", user.ID)})
		}
		if userCanDeleteUsers {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/user/delete/%d", user.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		row := ListingRow{Columns: &columns}
		listingRows = append(listingRows, &row)
	}
	return &UserListResponse{
		Response: NewResponse(headerText, currentUser, headerContent),
		Listing: &Listing{
			Header: listingHeader,
			Rows:   &listingRows,
		},
	}
}
