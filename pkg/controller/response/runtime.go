package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// RuntimeResponse is the struct for the runtime page.
type RuntimeResponse struct {
	*Response
	Runtime *model.Runtime
}

// RuntimeDetailResponse is the struct for the runtime detail page.
type RuntimeDetailResponse struct {
	*RuntimeResponse
	Details *DetailItems
}

// NewRuntimeDetailResponse is a constructor for the RuntimeDetailResponse struct.
func NewRuntimeDetailResponse(currentUser *model.User, runtime *model.Runtime) *RuntimeDetailResponse {
	header := &HeaderBlock{
		Title:       "Runtime Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/runtime/update/%d", runtime.ID),
				Privilege: "runtimes.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/runtime/delete/%d", runtime.ID),
				Privilege: "runtimes.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/runtime/list",
				Privilege: "runtimes.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", runtime.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: runtime.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: runtime.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: runtime.UpdatedAt}}},
	}
	return &RuntimeDetailResponse{
		RuntimeResponse: &RuntimeResponse{
			Response: NewResponse("Runtime Detail", currentUser, header),
			Runtime:  runtime,
		},
		Details: details,
	}
}

// RuntimeFormResponse is the struct for the runtime form responses.
type RuntimeFormResponse struct {
	*RuntimeDetailResponse
	FormItems []*FormItem
}

// NewRuntimeFormResponse is a constructor for the RuntimeFormResponse struct.
func NewRuntimeFormResponse(title string, currentUser *model.User, runtime *model.Runtime) *RuntimeFormResponse {
	runtimeDetailResponse := NewRuntimeDetailResponse(currentUser, runtime)
	runtimeDetailResponse.Header.Title = title
	runtimeDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	runtimeDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/runtime/list"), Privilege: "runtimes.view"}}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: runtime.Name, Required: true},
	}
	return &RuntimeFormResponse{
		RuntimeDetailResponse: runtimeDetailResponse,
		FormItems:             formItems,
	}
}

// RuntimeListResponse is the struct for the runtime list page.
type RuntimeListResponse struct {
	*Response
	Runtimes *model.Runtimes
}

// NewRuntimeListResponse is a constructor for the RuntimeListResponse struct.
func NewRuntimeListResponse(currentUser *model.User, runtimes *model.Runtimes) *RuntimeListResponse {
	header := &HeaderBlock{
		Title:       "Runtime List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/runtime/create",
				Privilege: "runtimes.create",
			},
		},
	}
	return &RuntimeListResponse{
		Response: NewResponse("Runtime List", currentUser, header),
		Runtimes: runtimes,
	}
}
