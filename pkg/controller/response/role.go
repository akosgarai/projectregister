package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewRoleDetailResponse is a constructor for the DetailResponse struct for a pool.
func NewRoleDetailResponse(currentUser *model.User, role *model.Role) *DetailResponse {
	headerText := "Role Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "roles", fmt.Sprintf("%d", role.ID)))
	resourceValues := components.DetailValues{}
	if len(role.Resources) > 0 {
		for _, resource := range role.Resources {
			resourceValues = append(resourceValues, &components.DetailValue{Value: resource.Name})
		}
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", role.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: role.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: role.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: role.UpdatedAt}}},
		{Label: "Resources", Value: &resourceValues},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateRoleResponse is a constructor for the FormResponse struct for creating a new role.
func NewCreateRoleResponse(currentUser *model.User, resources *model.Resources) *FormResponse {
	return newRoleFormResponse("Create Role", currentUser, &model.Role{}, resources, "/admin/role/create", "POST", "Create")
}

// NewUpdateRoleResponse is a constructor for the FormResponse struct for updating a role.
func NewUpdateRoleResponse(currentUser *model.User, role *model.Role, resources *model.Resources) *FormResponse {
	return newRoleFormResponse("Update Role", currentUser, role, resources, fmt.Sprintf("/admin/role/update/%d", role.ID), "POST", "Update")
}

// newRoleFormResponse is a constructor for the FormResponse struct for a role.
func newRoleFormResponse(title string, currentUser *model.User, role *model.Role, resources *model.Resources, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/role/list")})
	selectedOptions := []int64{}
	for _, resource := range role.Resources {
		if role.HasResource(resource.Name) {
			selectedOptions = append(selectedOptions, resource.ID)
		}
	}
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", role.Name, true, nil, nil),
		// Resources.
		components.NewFormItem("Resources", "resources", "checkboxgroup", "", true, resources.ToMap(), selectedOptions),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewRoleListResponse is a constructor for the ListingResponse struct of the roles.
func NewRoleListResponse(currentUser *model.User, roles *model.Roles, filter *model.RoleFilter) *ListingResponse {
	headerText := "Role List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("roles.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/role/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("roles.update")
	userCanDelete := currentUser.HasPrivilege("roles.delete")
	for _, role := range *roles {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: fmt.Sprintf("%d", role.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{Values: &components.ListingColumnValues{{Value: role.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{Values: &components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/role/view/%d", role.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/role/update/%d", role.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/role/delete/%d", role.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	/* Create the search form. The only form item is the name. */
	formItems := []*components.FormItem{
		components.NewFormItem("Name", "name", "text", filter.Name, false, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/role/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
