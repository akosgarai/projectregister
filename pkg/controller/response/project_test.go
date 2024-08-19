package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewProjectDetailResponse is a test function for the NewProjectDetailResponse function.
// It tests the response generation.
func TestNewProjectDetailResponse(t *testing.T) {
	project := &model.Project{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"projects.view"})
	response := NewProjectDetailResponse(testUser, project)
	if response.Title != "Project Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Project Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateProjectResponse is a test function for the NewCreateProjectResponse function.
// It tests the response generation.
func TestNewCreateProjectResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"projects.view"})
	response := NewCreateProjectResponse(testUser)
	if response.Title != "Create Project" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Project" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateProjectResponse is a test function for the NewUpdateProjectResponse function.
// It tests the response generation.
func TestNewUpdateProjectResponse(t *testing.T) {
	project := &model.Project{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"projects.view"})
	response := NewUpdateProjectResponse(testUser, project)
	if response.Title != "Update Project" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Project" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewProjectListResponse is a test function for the NewProjectListResponse function.
// It tests the response generation.
func TestNewProjectListResponse(t *testing.T) {
	projects := &model.Projects{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"projects.view", "projects.create", "projects.update", "projects.delete"})
	response := NewProjectListResponse(testUser, projects)
	if response.Title != "Project List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Project List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
