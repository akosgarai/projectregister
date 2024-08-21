package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewEnvironmentDetailResponse is a test function for the NewEnvironmentDetailResponse function.
// It tests the response generation.
func TestNewEnvironmentDetailResponse(t *testing.T) {
	environment := &model.Environment{
		ID:        1,
		Name:      "test environment",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
		Servers:   model.Servers{{ID: 1, Name: "test server"}},
		Databases: model.Databases{{ID: 1, Name: "test database"}},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"environments.view", "applications.create"})
	response := NewEnvironmentDetailResponse(testUser, environment)
	if response.Title != "Environment Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Environment Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 7 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateEnvironmentResponse is a test function for the NewCreateEnvironmentResponse function.
// It tests the response generation.
func TestNewCreateEnvironmentResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"environments.view"})
	servers := model.Servers{{ID: 1, Name: "test server"}}
	databases := model.Databases{{ID: 1, Name: "test database"}}
	response := NewCreateEnvironmentResponse(testUser, &servers, &databases)
	if response.Title != "Create Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 4 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateEnvironmentResponse is a test function for the NewUpdateEnvironmentResponse function.
// It tests the response generation.
func TestNewUpdateEnvironmentResponse(t *testing.T) {
	servers := model.Servers{{ID: 1, Name: "test server"}}
	databases := model.Databases{{ID: 1, Name: "test database"}}
	environment := &model.Environment{
		ID:          1,
		Name:        "test environment",
		Description: "test description",
		Servers:     servers,
		Databases:   databases,
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"environments.view"})
	response := NewUpdateEnvironmentResponse(testUser, environment, &servers, &databases)
	if response.Title != "Update Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 4 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewEnvironmentListResponse is a test function for the NewEnvironmentListResponse function.
// It tests the response generation.
func TestNewEnvironmentListResponse(t *testing.T) {
	environments := &model.Environments{
		{ID: 1, Name: "test environment", Description: "nope", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2 environment", Description: "nope", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"environments.view", "environments.create", "environments.update", "environments.delete"})
	response := NewEnvironmentListResponse(testUser, environments)
	if response.Title != "Environment List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Environment List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
