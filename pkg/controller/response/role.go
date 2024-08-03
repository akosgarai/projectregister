package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewRoleDetailResponse is a constructor for the DetailResponse struct for a pool.
func NewRoleDetailResponse(currentUser *model.User, role *model.Role) *DetailResponse {
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

// RoleFormResponse is the struct for the role form responses.
type RoleFormResponse struct {
	*DetailResponse
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
		DetailResponse: roleDetailResponse,
		FormItems:      formItems,
	}
}

// NewRoleListResponse is a constructor for the ListingResponse struct of the roles.
func NewRoleListResponse(currentUser *model.User, roles *model.Roles) *ListingResponse {
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
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", role.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: role.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
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

	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
