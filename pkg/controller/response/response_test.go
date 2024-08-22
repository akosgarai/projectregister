package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/resources"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewResponse is a test function for the NewResponse function.
// It tests the side menu generation.
func TestNewResponse(t *testing.T) {
	title := "title"
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	header := components.NewContentHeader("header", []*components.Link{})
	response := NewResponse(title, testUser, header)
	if response.Title != title {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header != header {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.SideMenu) != 1 {
		t.Errorf("Side menu is not generated properly. Got: %v", response.SideMenu)
	}
}

// TestNewDetailResponse is a test function for the NewDetailResponse function.
// It tests the response generation.
func TestNewDetailResponse(t *testing.T) {
	title := "title"
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	header := components.NewContentHeader("header", []*components.Link{})
	details := components.DetailItems{}
	response := NewDetailResponse(title, testUser, header, &details)
	if response.Title != title {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header != header {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if response.Details != &details {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewDetailHeaderButtons is a test function for the newDetailHeaderButtons function.
// It tests the button generation.
func TestNewDetailHeaderButtons(t *testing.T) {
	// test all the resources.
	for _, resource := range resources.Resources {
		// User with view priv only allowed to have the list button.
		testUser := testhelper.GetUserWithAccessToResources(1, []string{resource + ".view"})
		buttons := newDetailHeaderButtons(testUser, resource, "1")
		resourceSingular := resource[:len(resource)-1]
		if len(buttons) != 1 {
			t.Errorf("Buttons are not generated properly. Got: %v", buttons)
		}
		// The button has to be the list button link to the listing.
		if buttons[0].Text != "List" || buttons[0].Href != "/admin/"+resourceSingular+"/list" {
			t.Errorf("List button is not generated properly. Got: %s / %s", buttons[0].Text, buttons[0].Href)
		}
		// User with view and update priv allowed to have the list and update buttons.
		testUser = testhelper.GetUserWithAccessToResources(1, []string{resource + ".view", resource + ".update"})
		buttons = newDetailHeaderButtons(testUser, resource, "1")
		if len(buttons) != 2 {
			t.Errorf("Buttons are not generated properly. Got: %v", buttons)
		}
		// The buttons have to be the list and update buttons.
		// The first button has to be the update button links to the edit page.
		if buttons[0].Text != "Edit" || buttons[0].Href != "/admin/"+resourceSingular+"/update/1" {
			t.Errorf("Edit button is not generated properly. Got: %s / %s", buttons[0].Text, buttons[0].Href)
		}
		// The second button has to be the list button links to the listing.
		if buttons[1].Text != "List" || buttons[1].Href != "/admin/"+resourceSingular+"/list" {
			t.Errorf("Edit button is not generated properly. Got: %s / %s", buttons[1].Text, buttons[1].Href)
		}
		// User with view, update and delete priv allowed to have the list, update and delete buttons.
		testUser = testhelper.GetUserWithAccessToResources(1, []string{resource + ".view", resource + ".update", resource + ".delete"})
		buttons = newDetailHeaderButtons(testUser, resource, "1")
		if len(buttons) != 3 {
			t.Errorf("Buttons are not generated properly. Got: %v", buttons)
		}
		// The first button has to be the update button links to the edit page.
		if buttons[0].Text != "Edit" || buttons[0].Href != "/admin/"+resourceSingular+"/update/1" {
			t.Errorf("Edit button is not generated properly. Got: %s / %s", buttons[0].Text, buttons[0].Href)
		}
		// The second button has to be the delete button links to the delete page.
		if buttons[1].Text != "Delete" || buttons[1].Href != "/admin/"+resourceSingular+"/delete/1" {
			t.Errorf("Edit button is not generated properly. Got: %s / %s", buttons[1].Text, buttons[1].Href)
		}
		// The third button has to be the list button links to the listing.
		if buttons[2].Text != "List" || buttons[2].Href != "/admin/"+resourceSingular+"/list" {
			t.Errorf("Edit button is not generated properly. Got: %s / %s", buttons[2].Text, buttons[2].Href)
		}
	}
}

// TestNewListingResponse is a test function for the NewListingResponse function.
// It tests the response generation.
func TestNewListingResponse(t *testing.T) {
	title := "title"
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	header := components.NewContentHeader("header", []*components.Link{})
	listing := &components.Listing{
		Header: &components.ListingHeader{},
		Rows:   &components.ListingRows{},
	}
	response := NewListingResponse(title, testUser, header, listing)
	if response.Title != title {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header != header {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if response.Listing != listing {
		t.Errorf("Listing is not set properly. Got: %v", response.Listing)
	}
}

// TestNewFormResponse is a test function for the NewFormResponse function.
// It tests the response generation.
func TestNewFormResponse(t *testing.T) {
	title := "title"
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	header := components.NewContentHeader("header", []*components.Link{})
	var formItems []*components.FormItem
	form := &components.Form{
		Items: formItems,
	}
	response := NewFormResponse(title, testUser, header, form)
	if response.Title != title {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header != header {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if response.Form != form {
		t.Errorf("Form is not set properly. Got: %v", response.Form)
	}
}

// TestNewApplicationImportMappingResponse is a test function for the NewApplicationImportMappingResponse function.
// It tests the response generation.
func TestNewApplicationImportMappingResponse(t *testing.T) {
	title := "title"
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"users.view"})
	header := components.NewContentHeader("header", []*components.Link{})
	listing := &components.Listing{
		Header: &components.ListingHeader{},
		Rows:   &components.ListingRows{},
	}
	var formItems []*components.FormItem
	mappingForm := &components.Form{
		Items: formItems,
	}
	response := NewApplicationImportMappingResponse(title, testUser, header, listing, mappingForm)
	if response.Title != title {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header != header {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if response.Listing != listing {
		t.Errorf("Listing is not set properly. Got: %v", response.Listing)
	}
	if response.Form != mappingForm {
		t.Errorf("Form is not set properly. Got: %v", response.Form)
	}
}
