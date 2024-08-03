package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// PoolResponse is the struct for the pool page.
type PoolResponse struct {
	*Response
	Pool *model.Pool
}

// PoolDetailResponse is the struct for the pool detail page.
type PoolDetailResponse struct {
	*PoolResponse
	Details *DetailItems
}

// NewPoolDetailResponse is a constructor for the PoolDetailResponse struct.
func NewPoolDetailResponse(currentUser *model.User, pool *model.Pool) *PoolDetailResponse {
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
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", pool.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: pool.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: pool.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: pool.UpdatedAt}}},
	}
	return &PoolDetailResponse{
		PoolResponse: &PoolResponse{
			Response: NewResponse(headerText, currentUser, headerContent),
			Pool:     pool,
		},
		Details: details,
	}
}

// PoolFormResponse is the struct for the pool form responses.
type PoolFormResponse struct {
	*PoolDetailResponse
	FormItems []*FormItem
}

// NewPoolFormResponse is a constructor for the PoolFormResponse struct.
func NewPoolFormResponse(title string, currentUser *model.User, pool *model.Pool) *PoolFormResponse {
	poolDetailResponse := NewPoolDetailResponse(currentUser, pool)
	poolDetailResponse.Header.Title = title
	poolDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	poolDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/pool/list")}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: pool.Name, Required: true},
	}
	return &PoolFormResponse{
		PoolDetailResponse: poolDetailResponse,
		FormItems:          formItems,
	}
}

// PoolListResponse is the struct for the pool list page.
type PoolListResponse struct {
	*Response
	Listing *Listing
}

// NewPoolListResponse is a constructor for the PoolListResponse struct.
func NewPoolListResponse(currentUser *model.User, pools *model.Pools) *PoolListResponse {
	headerText := "Pool List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("pools.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/pool/create"))
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("pools.update")
	userCanDelete := currentUser.HasPrivilege("pools.delete")
	for _, pool := range *pools {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", pool.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: pool.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/pool/view/%d", pool.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/pool/update/%d", pool.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/pool/delete/%d", pool.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &PoolListResponse{
		Response: NewResponse(headerText, currentUser, headerContent),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
