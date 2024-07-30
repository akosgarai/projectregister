package response

import (
	"fmt"

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
	header := &HeaderBlock{
		Title:       "Pool Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/pool/update/%d", pool.ID),
				Privilege: "pools.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/pool/delete/%d", pool.ID),
				Privilege: "pools.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/pool/list",
				Privilege: "pools.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", pool.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: pool.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: pool.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: pool.UpdatedAt}}},
	}
	return &PoolDetailResponse{
		PoolResponse: &PoolResponse{
			Response: NewResponse("Pool Detail", currentUser, header),
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
	poolDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/pool/list"), Privilege: "pools.view"}}
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
	header := &HeaderBlock{
		Title:       "Pool List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/pool/create",
				Privilege: "pools.create",
			},
		},
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
			{Value: "View", Link: fmt.Sprintf("/admin/pool/detail/%d", pool.ID)},
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
		Response: NewResponse("Pool List", currentUser, header),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
