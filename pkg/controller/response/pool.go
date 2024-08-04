package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewPoolDetailResponse is a constructor for the DetailResponse struct for a pool.
func NewPoolDetailResponse(currentUser *model.User, pool *model.Pool) *DetailResponse {
	headerText := "Pool Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("pools.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/pool/update/%d", pool.ID)))
	}
	if currentUser.HasPrivilege("pools.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/pool/delete/%d", pool.ID)))
	}
	if currentUser.HasPrivilege("pools.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/pool/list"))
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", pool.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: pool.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: pool.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: pool.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreatePoolResponse is a constructor for the FormResponse struct for creating a new pool.
func NewCreatePoolResponse(currentUser *model.User) *FormResponse {
	return newPoolFormResponse("Create Pool", currentUser, &model.Pool{}, "/admin/pool/create", "POST", "Create")
}

// NewUpdatePoolResponse is a constructor for the FormResponse struct for updating a pool.
func NewUpdatePoolResponse(currentUser *model.User, pool *model.Pool) *FormResponse {
	return newPoolFormResponse("Update Pool", currentUser, pool, fmt.Sprintf("/admin/pool/update/%d", pool.ID), "POST", "Update")
}

// newPoolFormResponse is a constructor for the FormResponse struct for a pool.
func newPoolFormResponse(title string, currentUser *model.User, pool *model.Pool, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/pool/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", pool.Name, true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewPoolListResponse is a constructor for the ListingResponse struct of the pools.
func NewPoolListResponse(currentUser *model.User, pools *model.Pools) *ListingResponse {
	headerText := "Pool List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("pools.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/pool/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("pools.update")
	userCanDelete := currentUser.HasPrivilege("pools.delete")
	for _, pool := range *pools {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", pool.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: pool.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/pool/view/%d", pool.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/pool/update/%d", pool.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/pool/delete/%d", pool.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
