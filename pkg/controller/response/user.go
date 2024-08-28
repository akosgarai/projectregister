package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/transformers"
)

// NewUserDetailResponse is a constructor for the DetailResponse struct for a user.
func NewUserDetailResponse(currentUser, user *model.User) *DetailResponse {
	headerText := "User Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "users", fmt.Sprintf("%d", user.ID)))
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

// NewCreateUserResponse is a constructor for the FormResponse struct for the user create page.
func NewCreateUserResponse(currentUser *model.User, roles *model.Roles) *FormResponse {
	return newUserFormResponse("Create User", currentUser, &model.User{}, roles, "/admin/user/create", "POST", "Create")
}

// NewUpdateUserResponse is a constructor for the FormResponse struct for the user update page.
func NewUpdateUserResponse(currentUser, user *model.User, roles *model.Roles) *FormResponse {
	return newUserFormResponse("Update User", currentUser, user, roles, fmt.Sprintf("/admin/user/update/%d", user.ID), "POST", "Update")
}

// newUserFormResponse is a constructor for the FormResponse struct for a user.
func newUserFormResponse(title string, currentUser, user *model.User, roles *model.Roles, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/user/list")})
	roleID := ""
	selectedOptions := []int64{0}
	if user.Role != nil && user.Role.ID > 0 {
		roleID = fmt.Sprintf("%d", user.Role.ID)
		selectedOptions[0] = user.Role.ID
	}
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", user.Name, true, nil, nil),
		// Email.
		components.NewFormItem("Email", "email", "email", user.Email, true, nil, nil),
		// Password.
		components.NewFormItem("Password", "password", "password", "", false, nil, nil),
		// Roles.
		components.NewFormItem("Role", "role", "select", roleID, true, roles.ToMap(), selectedOptions),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewUserListResponse is a constructor for the ListingResponse struct of the users.
func NewUserListResponse(currentUser *model.User, users []*model.User, roles *model.Roles, filter *model.UserFilter) *ListingResponse {
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
		idColumn := components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", user.ID)}}}
		columns = append(columns, &idColumn)
		// Name
		nameColumn := components.ListingColumn{Values: &components.ListingColumnValues{{Value: user.Name}}}
		columns = append(columns, &nameColumn)
		// Email
		emailColumn := components.ListingColumn{Values: &components.ListingColumnValues{{Value: user.Email}}}
		columns = append(columns, &emailColumn)
		// Role
		roleLink := ""
		if userCanViewRoles {
			roleLink = fmt.Sprintf("/admin/role/view/%d", user.Role.ID)
		}
		roleColumn := components.ListingColumn{Values: &components.ListingColumnValues{{Value: user.Role.Name, Link: roleLink}}}
		// Actions
		columns = append(columns, &roleColumn)
		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
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
	formItems := []*components.FormItem{
		components.NewFormItem("Name", "name", "text", filter.Name, false, nil, nil),
		components.NewFormItem("Email", "email", "text", filter.Email, false, nil, nil),
		// Roles.
		components.NewFormItem("Role", "role", "multiselect", "", false, roles.ToMap(), transformers.StringSliceToInt64Slice(filter.RoleIDs)),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/user/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
