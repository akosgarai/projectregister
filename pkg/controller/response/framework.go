package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewFrameworkDetailResponse is a constructor for the DetailResponse struct for a framework.
func NewFrameworkDetailResponse(currentUser *model.User, framework *model.Framework) *DetailResponse {
	headerText := "Framework Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "frameworks", fmt.Sprintf("%d", framework.ID)))
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", framework.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: framework.Name}}},
		{Label: "Score", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", framework.Score)}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: framework.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: framework.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateFrameworkResponse is a constructor for the FormResponse struct for creating a framework.
func NewCreateFrameworkResponse(currentUser *model.User) *FormResponse {
	return newFrameworkFormResponse("Create Framework", currentUser, &model.Framework{}, "/admin/framework/create", "POST", "Create")
}

// NewUpdateFrameworkResponse is a constructor for the FormResponse struct for updating a framework.
func NewUpdateFrameworkResponse(currentUser *model.User, framework *model.Framework) *FormResponse {
	return newFrameworkFormResponse("Update Framework", currentUser, framework, fmt.Sprintf("/admin/framework/update/%d", framework.ID), "POST", "Update")
}

// newFrameworkFormResponse is a constructor for the FormResponse struct for a framework.
func newFrameworkFormResponse(title string, currentUser *model.User, framework *model.Framework, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/framework/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", framework.Name, true, nil, nil),
		// Score.
		components.NewFormItem("Score", "score", "number", fmt.Sprintf("%d", framework.Score), true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewFrameworkListResponse is a constructor for the ListingResponse struct of the frameworks.
func NewFrameworkListResponse(currentUser *model.User, frameworks *model.Frameworks, filter *model.FrameworkFilter) *ListingResponse {
	headerText := "Framework List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("frameworks.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/framework/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Score", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("frameworks.update")
	userCanDelete := currentUser.HasPrivilege("frameworks.delete")
	for _, framework := range *frameworks {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", framework.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: framework.Name}}}
		columns = append(columns, nameColumn)
		scoreColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", framework.Score)}}}
		columns = append(columns, scoreColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/framework/view/%d", framework.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/framework/update/%d", framework.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/framework/delete/%d", framework.ID), Form: true})
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
		Action: "/admin/framework/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
