package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewClientDetailResponse is a test function for the NewClientDetailResponse function.
// It tests the response generation.
func TestNewClientDetailResponse(t *testing.T) {
	client := &model.Client{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"clients.view"})
	response := NewClientDetailResponse(testUser, client)
	if response.Title != "Client Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Client Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateClientResponse is a test function for the NewCreateClientResponse function.
// It tests the response generation.
func TestNewCreateClientResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"clients.view"})
	response := NewCreateClientResponse(testUser)
	if response.Title != "Create Client" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Client" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateClientResponse is a test function for the NewUpdateClientResponse function.
// It tests the response generation.
func TestNewUpdateClientResponse(t *testing.T) {
	client := &model.Client{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"clients.view"})
	response := NewUpdateClientResponse(testUser, client)
	if response.Title != "Update Client" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Client" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewClientListResponse is a test function for the NewClientListResponse function.
// It tests the response generation.
func TestNewClientListResponse(t *testing.T) {
	clients := &model.Clients{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"clients.view", "clients.create", "clients.update", "clients.delete"})
	response := NewClientListResponse(testUser, clients)
	if response.Title != "Client List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Client List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
