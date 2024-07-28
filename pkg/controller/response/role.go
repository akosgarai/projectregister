package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// RoleDetailResponse is the struct for the role detail page.
type RoleDetailResponse struct {
	*Response
	Header *HeaderBlock
	Role   *model.Role
}

// NewRoleDetailResponse is a constructor for the RoleDetailResponse struct.
func NewRoleDetailResponse(currentUser *model.User, role *model.Role) *RoleDetailResponse {
	return &RoleDetailResponse{
		Response: NewResponse("Role Detail", currentUser),
		Header: &HeaderBlock{
			Title:       "Role Detail",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Edit",
					Link:      fmt.Sprintf("/admin/role/update/%d", role.ID),
					Privilege: "roles.update",
				},
				{
					Label:     "Delete",
					Link:      fmt.Sprintf("/admin/role/delete/%d", role.ID),
					Privilege: "roles.delete",
				},
			},
		},
		Role: role,
	}
}

// RoleFormResponse is the struct for the role form responses.
type RoleFormResponse struct {
	*RoleDetailResponse
	Resources []*model.Resource
}

// NewRoleFormResponse is a constructor for the RoleFormResponse struct.
func NewRoleFormResponse(title string, currentUser *model.User, role *model.Role, resources []*model.Resource) *RoleFormResponse {
	roleDetailResponse := NewRoleDetailResponse(currentUser, role)
	roleDetailResponse.Header.Title = title
	roleDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	roleDetailResponse.Header.Buttons = []*ActionButton{}
	return &RoleFormResponse{
		RoleDetailResponse: roleDetailResponse,
		Resources:          resources,
	}
}

// RoleListResponse is the struct for the role list page.
type RoleListResponse struct {
	*Response
	Header *HeaderBlock
	Roles  []*model.Role
}

// NewRoleListResponse is a constructor for the RoleListResponse struct.
func NewRoleListResponse(currentUser *model.User, roles []*model.Role) *RoleListResponse {
	return &RoleListResponse{
		Response: NewResponse("Role List", currentUser),
		Header: &HeaderBlock{
			Title:       "Role List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/role/create",
					Privilege: "roles.create",
				},
			},
		},
		Roles: roles,
	}
}
