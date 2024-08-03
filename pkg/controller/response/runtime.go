package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
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
	headerText := "Runtime Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("runtimes.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/runtime/update/%d", runtime.ID)))
	}
	if currentUser.HasPrivilege("runtimes.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/runtime/delete/%d", runtime.ID)))
	}
	if currentUser.HasPrivilege("runtimes.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/runtime/list"))
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", runtime.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: runtime.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: runtime.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: runtime.UpdatedAt}}},
	}
	return &RuntimeDetailResponse{
		RuntimeResponse: &RuntimeResponse{
			Response: NewResponse(headerText, currentUser, headerContent),
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
	runtimeDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/runtime/list")}
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
	Listing *Listing
}

// NewRuntimeListResponse is a constructor for the RuntimeListResponse struct.
func NewRuntimeListResponse(currentUser *model.User, runtimes *model.Runtimes) *RuntimeListResponse {
	headerText := "Runtime List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("runtimes.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/runtime/create"))
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("runtimes.update")
	userCanDelete := currentUser.HasPrivilege("runtimes.delete")
	for _, runtime := range *runtimes {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", runtime.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: runtime.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/runtime/view/%d", runtime.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/runtime/update/%d", runtime.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/runtime/delete/%d", runtime.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &RuntimeListResponse{
		Response: NewResponse(headerText, currentUser, headerContent),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
