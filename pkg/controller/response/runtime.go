package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewRuntimeDetailResponse is a constructor for the DetailResponse struct for a runtime.
func NewRuntimeDetailResponse(currentUser *model.User, runtime *model.Runtime) *DetailResponse {
	headerText := "Runtime Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "runtimes", fmt.Sprintf("%d", runtime.ID)))
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", runtime.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: runtime.Name}}},
		{Label: "Score", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", runtime.Score)}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: runtime.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: runtime.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateRuntimeResponse is a constructor for the FormResponse struct for creating a runtime.
func NewCreateRuntimeResponse(currentUser *model.User) *FormResponse {
	return newRuntimeFormResponse("Create Runtime", currentUser, &model.Runtime{}, "/admin/runtime/create", "POST", "Create")
}

// NewUpdateRuntimeResponse is a constructor for the FormResponse struct for updating a runtime.
func NewUpdateRuntimeResponse(currentUser *model.User, runtime *model.Runtime) *FormResponse {
	return newRuntimeFormResponse("Update Runtime", currentUser, runtime, fmt.Sprintf("/admin/runtime/update/%d", runtime.ID), "POST", "Update")
}

// newRuntimeFormResponse is a constructor for the FormResponse struct for a runtime.
func newRuntimeFormResponse(title string, currentUser *model.User, runtime *model.Runtime, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/runtime/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", runtime.Name, true, nil, nil),
		// Score.
		components.NewFormItem("Score", "score", "number", fmt.Sprintf("%d", runtime.Score), true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewRuntimeListResponse is a constructor for the ListingResponse struct of the runtimes.
func NewRuntimeListResponse(currentUser *model.User, runtimes *model.Runtimes, filter *model.RuntimeFilter) *ListingResponse {
	headerText := "Runtime List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("runtimes.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/runtime/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Score", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("runtimes.update")
	userCanDelete := currentUser.HasPrivilege("runtimes.delete")
	for _, runtime := range *runtimes {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", runtime.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: runtime.Name}}}
		columns = append(columns, nameColumn)
		scoreColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", runtime.Score)}}}
		columns = append(columns, scoreColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/runtime/view/%d", runtime.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/runtime/update/%d", runtime.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/runtime/delete/%d", runtime.ID), Form: true})
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
		Action: "/admin/runtime/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
