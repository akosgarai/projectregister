package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewRoleDetailResponse is a test function for the NewRoleDetailResponse function.
// It tests the response generation.
func TestNewRoleDetailResponse(t *testing.T) {
	role := &model.Role{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
		Resources: model.Resources{
			{ID: 1, Name: "roles.view"},
			{ID: 2, Name: "roles.update"},
		},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"roles.view"})
	response := NewRoleDetailResponse(testUser, role)
	if response.Title != "Role Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Role Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 5 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateRoleResponse is a test function for the NewCreateRoleResponse function.
// It tests the response generation.
func TestNewCreateRoleResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"roles.view"})
	response := NewCreateRoleResponse(testUser, &testUser.Role.Resources)
	if response.Title != "Create Role" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Role" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 2 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateRoleResponse is a test function for the NewUpdateRoleResponse function.
// It tests the response generation.
func TestNewUpdateRoleResponse(t *testing.T) {
	role := &model.Role{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
		Resources: model.Resources{
			{ID: 1, Name: "roles.view"},
			{ID: 2, Name: "roles.update"},
		},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"roles.view"})
	response := NewUpdateRoleResponse(testUser, role, &testUser.Role.Resources)
	if response.Title != "Update Role" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Role" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 2 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewRoleListResponse is a test function for the NewRoleListResponse function.
// It tests the response generation.
func TestNewRoleListResponse(t *testing.T) {
	roles := &model.Roles{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01", Resources: model.Resources{{ID: 1, Name: "roles.view"}, {ID: 2, Name: "roles.update"}}},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01", Resources: model.Resources{{ID: 1, Name: "roles.view"}, {ID: 2, Name: "roles.update"}}},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"roles.view", "roles.create", "roles.update", "roles.delete"})
	response := NewRoleListResponse(testUser, roles)
	if response.Title != "Role List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Role List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
