package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewUserDetailResponse is a test function for the NewUserDetailResponse function.
// It tests the response generation.
func TestNewUserDetailResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "roles.view"})
	user := testhelper.GetUserWithAccessToResources(2, []string{"users.view"})
	response := NewUserDetailResponse(testUser, user)
	if response.Title != "User Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "User Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %d", len(*response.Details))
	}
}

// TestNewCreateUserResponse is a test function for the NewCreateUserResponse function.
// It tests the response generation.
func TestNewCreateUserResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	roles := &model.Roles{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	response := NewCreateUserResponse(testUser, roles)
	if response.Title != "Create User" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create User" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 4 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateUserResponse is a test function for the NewUpdateUserResponse function.
// It tests the response generation.
func TestNewUpdateUserResponse(t *testing.T) {
	user := testhelper.GetUserWithAccessToResources(2, []string{"users.view"})
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	roles := &model.Roles{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	response := NewUpdateUserResponse(testUser, user, roles)
	if response.Title != "Update User" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update User" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 4 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUserListResponse is a test function for the NewUserListResponse function.
// It tests the response generation.
func TestNewUserListResponse(t *testing.T) {
	users := []*model.User{
		testhelper.GetUserWithAccessToResources(2, []string{"users.view"}),
		testhelper.GetUserWithAccessToResources(3, []string{"users.view"}),
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view", "users.create", "users.update", "users.delete", "roles.view"})
	response := NewUserListResponse(testUser, users)
	if response.Title != "User List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "User List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
