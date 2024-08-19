package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewDatabaseDetailResponse is a test function for the NewDatabaseDetailResponse function.
// It tests the response generation.
func TestNewDatabaseDetailResponse(t *testing.T) {
	database := &model.Database{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"databases.view"})
	response := NewDatabaseDetailResponse(testUser, database)
	if response.Title != "Database Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Database Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateDatabaseResponse is a test function for the NewCreateDatabaseResponse function.
// It tests the response generation.
func TestNewCreateDatabaseResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"databases.view"})
	response := NewCreateDatabaseResponse(testUser)
	if response.Title != "Create Database" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Database" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateDatabaseResponse is a test function for the NewUpdateDatabaseResponse function.
// It tests the response generation.
func TestNewUpdateDatabaseResponse(t *testing.T) {
	database := &model.Database{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"databases.view"})
	response := NewUpdateDatabaseResponse(testUser, database)
	if response.Title != "Update Database" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Database" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewDatabaseListResponse is a test function for the NewDatabaseListResponse function.
// It tests the response generation.
func TestNewDatabaseListResponse(t *testing.T) {
	databases := &model.Databases{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"databases.view", "databases.create", "databases.update", "databases.delete"})
	response := NewDatabaseListResponse(testUser, databases)
	if response.Title != "Database List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Database List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
