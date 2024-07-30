package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// UserResponse is the struct for the user page.
type UserResponse struct {
	*Response
	User *model.User
}

// UserDetailResponse is the struct for the user detail page.
type UserDetailResponse struct {
	*UserResponse
	Details *DetailItems
}

// NewUserDetailResponse is a constructor for the UserDetailResponse struct.
func NewUserDetailResponse(currentUser, user *model.User) *UserDetailResponse {
	header := &HeaderBlock{
		Title:       "User Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/user/update/%d", user.ID),
				Privilege: "users.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/user/delete/%d", user.ID),
				Privilege: "users.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/user/list",
				Privilege: "users.view",
			},
		},
	}
	roleLink := ""
	if currentUser.HasPrivilege("roles.view") {
		roleLink = fmt.Sprintf("/admin/role/view/%d", user.Role.ID)
	}
	roleValue := DetailValues{{Value: user.Role.Name, Link: roleLink}}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", user.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: user.Name}}},
		{Label: "Email", Value: &DetailValues{{Value: user.Email}}},
		{Label: "Role", Value: &roleValue},
	}
	return &UserDetailResponse{
		UserResponse: &UserResponse{
			Response: NewResponse("User Detail", currentUser, header),
			User:     user,
		},
		Details: details,
	}
}

// UserFormResponse is the struct for the user form responses.
type UserFormResponse struct {
	*UserDetailResponse
	FormItems []*FormItem
}

// NewUserFormResponse is a constructor for the UserFormResponse struct.
func NewUserFormResponse(title string, currentUser, user *model.User, roles *model.Roles) *UserFormResponse {
	userDetailResponse := NewUserDetailResponse(currentUser, user)
	userDetailResponse.Header.Title = title
	userDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	userDetailResponse.Header.Buttons = []*ActionButton{{Label: "Back", Link: "/admin/user/list", Privilege: "users.view"}}
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
		UserDetailResponse: userDetailResponse,
		FormItems:          formItems,
	}
}

// UserListResponse is the struct for the user list page.
type UserListResponse struct {
	*Response
	Listing *Listing
}

// NewUserListResponse is a constructor for the UserListResponse struct.
func NewUserListResponse(currentUser *model.User, users []*model.User) *UserListResponse {
	header := &HeaderBlock{
		Title:       "User List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "New",
				Link:      "/admin/user/create",
				Privilege: "users.create",
			},
		},
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
		Response: NewResponse("User List", currentUser, header),
		Listing: &Listing{
			Header: listingHeader,
			Rows:   &listingRows,
		},
	}
}
