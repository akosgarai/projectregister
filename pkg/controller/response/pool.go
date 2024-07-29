package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// PoolDetailResponse is the struct for the pool detail page.
type PoolDetailResponse struct {
	*Response
	Pool *model.Pool
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
		},
	}
	return &PoolDetailResponse{
		Response: NewResponse("Pool Detail", currentUser, header),
		Pool:     pool,
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
	Pools *model.Pools
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
	return &PoolListResponse{
		Response: NewResponse("Pool List", currentUser, header),
		Pools:    pools,
	}
}
