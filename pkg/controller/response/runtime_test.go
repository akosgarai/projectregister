package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewRuntimeDetailResponse is a test function for the NewRuntimeDetailResponse function.
// It tests the response generation.
func TestNewRuntimeDetailResponse(t *testing.T) {
	runtime := &model.Runtime{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"runtimes.view"})
	response := NewRuntimeDetailResponse(testUser, runtime)
	if response.Title != "Runtime Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Runtime Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateRuntimeResponse is a test function for the NewCreateRuntimeResponse function.
// It tests the response generation.
func TestNewCreateRuntimeResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"runtimes.view"})
	response := NewCreateRuntimeResponse(testUser)
	if response.Title != "Create Runtime" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Runtime" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateRuntimeResponse is a test function for the NewUpdateRuntimeResponse function.
// It tests the response generation.
func TestNewUpdateRuntimeResponse(t *testing.T) {
	runtime := &model.Runtime{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"runtimes.view"})
	response := NewUpdateRuntimeResponse(testUser, runtime)
	if response.Title != "Update Runtime" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Runtime" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewRuntimeListResponse is a test function for the NewRuntimeListResponse function.
// It tests the response generation.
func TestNewRuntimeListResponse(t *testing.T) {
	runtimes := &model.Runtimes{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"runtimes.view", "runtimes.create", "runtimes.update", "runtimes.delete"})
	response := NewRuntimeListResponse(testUser, runtimes)
	if response.Title != "Runtime List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Runtime List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
