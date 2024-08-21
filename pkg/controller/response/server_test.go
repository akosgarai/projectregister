package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewServerDetailResponse is a test function for the NewServerDetailResponse function.
// It tests the response generation.
func TestNewServerDetailResponse(t *testing.T) {
	server := &model.Server{
		ID:        1,
		Name:      "test server",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
		Runtimes:  model.Runtimes{{ID: 1, Name: "test runtime"}},
		Pools:     model.Pools{{ID: 1, Name: "test pool"}},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"servers.view"})
	response := NewServerDetailResponse(testUser, server)
	if response.Title != "Server Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Server Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 8 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateServerResponse is a test function for the NewCreateServerResponse function.
// It tests the response generation.
func TestNewCreateServerResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"servers.view"})
	pools := model.Pools{{ID: 1, Name: "test pool"}}
	runtimes := model.Runtimes{{ID: 1, Name: "test runtime"}}
	response := NewCreateServerResponse(testUser, &pools, &runtimes)
	if response.Title != "Create Server" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Server" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 5 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateServerResponse is a test function for the NewUpdateServerResponse function.
// It tests the response generation.
func TestNewUpdateServerResponse(t *testing.T) {
	pools := model.Pools{{ID: 1, Name: "test pool"}}
	runtimes := model.Runtimes{{ID: 1, Name: "test runtime"}}
	server := &model.Server{
		ID:       1,
		Name:     "test",
		Runtimes: runtimes,
		Pools:    pools,
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"servers.view"})
	response := NewUpdateServerResponse(testUser, server, &pools, &runtimes)
	if response.Title != "Update Server" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Server" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 5 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewServerListResponse is a test function for the NewServerListResponse function.
// It tests the response generation.
func TestNewServerListResponse(t *testing.T) {
	servers := &model.Servers{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"servers.view", "servers.create", "servers.update", "servers.delete"})
	response := NewServerListResponse(testUser, servers)
	if response.Title != "Server List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Server List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
