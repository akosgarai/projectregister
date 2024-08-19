package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewPoolDetailResponse is a test function for the NewPoolDetailResponse function.
// It tests the response generation.
func TestNewPoolDetailResponse(t *testing.T) {
	pool := &model.Pool{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"pools.view"})
	response := NewPoolDetailResponse(testUser, pool)
	if response.Title != "Pool Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Pool Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreatePoolResponse is a test function for the NewCreatePoolResponse function.
// It tests the response generation.
func TestNewCreatePoolResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"pools.view"})
	response := NewCreatePoolResponse(testUser)
	if response.Title != "Create Pool" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Pool" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdatePoolResponse is a test function for the NewUpdatePoolResponse function.
// It tests the response generation.
func TestNewUpdatePoolResponse(t *testing.T) {
	pool := &model.Pool{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"pools.view"})
	response := NewUpdatePoolResponse(testUser, pool)
	if response.Title != "Update Pool" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Pool" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewPoolListResponse is a test function for the NewPoolListResponse function.
// It tests the response generation.
func TestNewPoolListResponse(t *testing.T) {
	pools := &model.Pools{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"pools.view", "pools.create", "pools.update", "pools.delete"})
	response := NewPoolListResponse(testUser, pools)
	if response.Title != "Pool List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Pool List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
