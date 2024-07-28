package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// RuntimeDetailResponse is the struct for the runtime detail page.
type RuntimeDetailResponse struct {
	*Response
	Header  *HeaderBlock
	Runtime *model.Runtime
}

// NewRuntimeDetailResponse is a constructor for the RuntimeDetailResponse struct.
func NewRuntimeDetailResponse(currentUser *model.User, runtime *model.Runtime) *RuntimeDetailResponse {
	return &RuntimeDetailResponse{
		Response: NewResponse("Runtime Detail", currentUser),
		Header: &HeaderBlock{
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
			},
		},
		Runtime: runtime,
	}
}

// RuntimeFormResponse is the struct for the runtime form responses.
type RuntimeFormResponse struct {
	*RuntimeDetailResponse
}

// NewRuntimeFormResponse is a constructor for the RuntimeFormResponse struct.
func NewRuntimeFormResponse(title string, currentUser *model.User, runtime *model.Runtime) *RuntimeFormResponse {
	runtimeDetailResponse := NewRuntimeDetailResponse(currentUser, runtime)
	runtimeDetailResponse.Header.Title = title
	runtimeDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	runtimeDetailResponse.Header.Buttons = []*ActionButton{}
	return &RuntimeFormResponse{
		RuntimeDetailResponse: runtimeDetailResponse,
	}
}

// RuntimeListResponse is the struct for the runtime list page.
type RuntimeListResponse struct {
	*Response
	Header   *HeaderBlock
	Runtimes []*model.Runtime
}

// NewRuntimeListResponse is a constructor for the RuntimeListResponse struct.
func NewRuntimeListResponse(currentUser *model.User, runtimes []*model.Runtime) *RuntimeListResponse {
	return &RuntimeListResponse{
		Response: NewResponse("Runtime List", currentUser),
		Header: &HeaderBlock{
			Title:       "Runtime List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/runtime/create",
					Privilege: "runtimes.create",
				},
			},
		},
		Runtimes: runtimes,
	}
}
