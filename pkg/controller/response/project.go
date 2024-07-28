package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// ProjectDetailResponse is the struct for the project detail page.
type ProjectDetailResponse struct {
	*Response
	Header  *HeaderBlock
	Project *model.Project
}

// NewProjectDetailResponse is a constructor for the ProjectDetailResponse struct.
func NewProjectDetailResponse(currentUser *model.User, project *model.Project) *ProjectDetailResponse {
	return &ProjectDetailResponse{
		Response: NewResponse("Project Detail", currentUser),
		Header: &HeaderBlock{
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
			},
		},
		Project: project,
	}
}

// ProjectFormResponse is the struct for the project form responses.
type ProjectFormResponse struct {
	*ProjectDetailResponse
}

// NewProjectFormResponse is a constructor for the ProjectFormResponse struct.
func NewProjectFormResponse(title string, currentUser *model.User, project *model.Project) *ProjectFormResponse {
	projectDetailResponse := NewProjectDetailResponse(currentUser, project)
	projectDetailResponse.Header.Title = title
	projectDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	projectDetailResponse.Header.Buttons = []*ActionButton{}
	return &ProjectFormResponse{
		ProjectDetailResponse: projectDetailResponse,
	}
}

// ProjectListResponse is the struct for the project list page.
type ProjectListResponse struct {
	*Response
	Header   *HeaderBlock
	Projects []*model.Project
}

// NewProjectListResponse is a constructor for the ProjectListResponse struct.
func NewProjectListResponse(currentUser *model.User, projects []*model.Project) *ProjectListResponse {
	return &ProjectListResponse{
		Response: NewResponse("Project List", currentUser),
		Header: &HeaderBlock{
			Title:       "Project List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/project/create",
					Privilege: "projects.create",
				},
			},
		},
		Projects: projects,
	}
}
