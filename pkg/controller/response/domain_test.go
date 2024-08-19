package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewDomainDetailResponse is a test function for the NewDomainDetailResponse function.
// It tests the response generation.
func TestNewDomainDetailResponse(t *testing.T) {
	domain := &model.Domain{
		ID:        1,
		Name:      "test",
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"domains.view"})
	response := NewDomainDetailResponse(testUser, domain)
	if response.Title != "Domain Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Domain Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 4 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateDomainResponse is a test function for the NewCreateDomainResponse function.
// It tests the response generation.
func TestNewCreateDomainResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"domains.view"})
	response := NewCreateDomainResponse(testUser)
	if response.Title != "Create Domain" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Domain" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateDomainResponse is a test function for the NewUpdateDomainResponse function.
// It tests the response generation.
func TestNewUpdateDomainResponse(t *testing.T) {
	domain := &model.Domain{
		ID:   1,
		Name: "test",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"domains.view"})
	response := NewUpdateDomainResponse(testUser, domain)
	if response.Title != "Update Domain" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Domain" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 1 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewDomainListResponse is a test function for the NewDomainListResponse function.
// It tests the response generation.
func TestNewDomainListResponse(t *testing.T) {
	domains := &model.Domains{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"domains.view", "domains.create", "domains.update", "domains.delete"})
	response := NewDomainListResponse(testUser, domains)
	if response.Title != "Domain List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Domain List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}
