package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ProjectResponse is the struct for the project page.
type ProjectResponse struct {
	*Response
	Project *model.Project
}

// ProjectDetailResponse is the struct for the project detail page.
type ProjectDetailResponse struct {
	*ProjectResponse
	Details *DetailItems
}

// NewProjectDetailResponse is a constructor for the ProjectDetailResponse struct.
func NewProjectDetailResponse(currentUser *model.User, project *model.Project) *ProjectDetailResponse {
	header := &HeaderBlock{
		Title:       "Project Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/project/update/%d", project.ID),
				Privilege: "projects.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/project/delete/%d", project.ID),
				Privilege: "projects.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/project/list",
				Privilege: "projects.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", project.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: project.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: project.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: project.UpdatedAt}}},
	}
	return &ProjectDetailResponse{
		ProjectResponse: &ProjectResponse{
			Response: NewResponse("Project Detail", currentUser, header),
			Project:  project,
		},
		Details: details,
	}
}

// ProjectFormResponse is the struct for the project form responses.
type ProjectFormResponse struct {
	*ProjectDetailResponse
	FormItems []*FormItem
}

// NewProjectFormResponse is a constructor for the ProjectFormResponse struct.
func NewProjectFormResponse(title string, currentUser *model.User, project *model.Project) *ProjectFormResponse {
	projectDetailResponse := NewProjectDetailResponse(currentUser, project)
	projectDetailResponse.Header.Title = title
	projectDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	projectDetailResponse.Header.Buttons = []*ActionButton{{Label: "List", Link: fmt.Sprintf("/admin/project/list"), Privilege: "projects.view"}}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: project.Name, Required: true},
	}
	return &ProjectFormResponse{
		ProjectDetailResponse: projectDetailResponse,
		FormItems:             formItems,
	}
}

// ProjectListResponse is the struct for the project list page.
type ProjectListResponse struct {
	*Response
	Projects *model.Projects
}

// NewProjectListResponse is a constructor for the ProjectListResponse struct.
func NewProjectListResponse(currentUser *model.User, projects *model.Projects) *ProjectListResponse {
	header := &HeaderBlock{
		Title:       "Project List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/project/create",
				Privilege: "projects.create",
			},
		},
	}
	return &ProjectListResponse{
		Response: NewResponse("Project List", currentUser, header),
		Projects: projects,
	}
}
