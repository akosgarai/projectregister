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
	Users []*model.User
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
	return &UserListResponse{
		Response: NewResponse("User List", currentUser, header),
		Users:    users,
	}
}
