package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewProjectDetailResponse is a constructor for the ProjectDetailResponse struct.
func NewProjectDetailResponse(currentUser *model.User, project *model.Project) *DetailResponse {
	headerText := "Project Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "projects", fmt.Sprintf("%d", project.ID)))
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", project.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: project.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: project.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: project.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateProjectResponse is a constructor for the FormResponse struct for the project create page.
func NewCreateProjectResponse(currentUser *model.User) *FormResponse {
	return newProjectFormResponse("Create Project", currentUser, &model.Project{}, "/admin/project/create", "POST", "Create")
}

// NewUpdateProjectResponse is a constructor for the FormResponse struct for the project update page.
func NewUpdateProjectResponse(currentUser *model.User, project *model.Project) *FormResponse {
	return newProjectFormResponse("Update Project", currentUser, project, fmt.Sprintf("/admin/project/update/%d", project.ID), "POST", "Update")
}

// NewProjectFormResponse is a constructor for the FormResponse struct for a project.
func newProjectFormResponse(title string, currentUser *model.User, project *model.Project, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/project/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", project.Name, true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewProjectListResponse is a constructor for the ListingResponse struct of the projects.
func NewProjectListResponse(currentUser *model.User, projects *model.Projects) *ListingResponse {
	headerText := "Project List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("projects.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/project/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("projects.update")
	userCanDelete := currentUser.HasPrivilege("projects.delete")
	for _, project := range *projects {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", project.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: project.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/project/view/%d", project.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/project/update/%d", project.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/project/delete/%d", project.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, nil)
}
