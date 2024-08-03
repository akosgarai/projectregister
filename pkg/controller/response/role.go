package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// RoleResponse is the base struct for the role responses.
type RoleResponse struct {
	*Response
	Role *model.Role
}

// RoleDetailResponse is the struct for the role detail page.
type RoleDetailResponse struct {
	*RoleResponse
	Details *DetailItems
}

// NewRoleDetailResponse is a constructor for the RoleDetailResponse struct.
func NewRoleDetailResponse(currentUser *model.User, role *model.Role) *RoleDetailResponse {
	headerText := "Role Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("roles.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/role/update/%d", role.ID)))
	}
	if currentUser.HasPrivilege("roles.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/role/delete/%d", role.ID)))
	}
	if currentUser.HasPrivilege("roles.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/role/list"))
	}
	resourceValues := DetailValues{}
	if len(role.Resources) > 0 {
		for _, resource := range role.Resources {
			resourceValues = append(resourceValues, &DetailValue{Value: resource.Name})
		}
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", role.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: role.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: role.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: role.UpdatedAt}}},
		{Label: "Resources", Value: &resourceValues},
	}
	return &RoleDetailResponse{
		RoleResponse: &RoleResponse{
			Response: NewResponse(headerText, currentUser, headerContent),
			Role:     role,
		},
		Details: details,
	}
}

// RoleFormResponse is the struct for the role form responses.
type RoleFormResponse struct {
	*RoleDetailResponse
	FormItems []*FormItem
}

// NewRoleFormResponse is a constructor for the RoleFormResponse struct.
func NewRoleFormResponse(title string, currentUser *model.User, role *model.Role, resources *model.Resources) *RoleFormResponse {
	roleDetailResponse := NewRoleDetailResponse(currentUser, role)
	roleDetailResponse.Header.Title = title
	roleDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	roleDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/role/list")}
	selectedOptions := SelectedOptions{}
	for _, resource := range role.Resources {
		if role.HasResource(resource.Name) {
			selectedOptions = append(selectedOptions, resource.ID)
		}
	}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: role.Name, Required: true},
		// Resources.
		{Label: "Resources", Type: "checkboxgroup", Name: "resources", Value: "", Required: true, Options: resources.ToMap(), SelectedOptions: selectedOptions},
	}
	return &RoleFormResponse{
		RoleDetailResponse: roleDetailResponse,
		FormItems:          formItems,
	}
}

// RoleListResponse is the struct for the role list page.
type RoleListResponse struct {
	*Response
	Listing *Listing
}

// NewRoleListResponse is a constructor for the RoleListResponse struct.
func NewRoleListResponse(currentUser *model.User, roles *model.Roles) *RoleListResponse {
	headerText := "Role List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("roles.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/role/create"))
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("roles.update")
	userCanDelete := currentUser.HasPrivilege("roles.delete")
	for _, role := range *roles {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", role.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: role.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/role/view/%d", role.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/role/update/%d", role.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/role/delete/%d", role.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}

	return &RoleListResponse{
		Response: NewResponse(headerText, currentUser, headerContent),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
