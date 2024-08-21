package response

import (
	"testing"

	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/parser"
	"github.com/akosgarai/projectregister/pkg/testhelper"
)

// TestNewApplicationDetailResponse is a test function for the NewApplicationDetailResponse function.
// It tests the response generation.
func TestNewApplicationDetailResponse(t *testing.T) {
	application := &model.Application{
		ID:           1,
		Client:       &model.Client{ID: 1, Name: "test client"},
		Project:      &model.Project{ID: 1, Name: "test project"},
		Environment:  &model.Environment{ID: 1, Name: "test environment"},
		Database:     &model.Database{ID: 1, Name: "test database"},
		DBName:       "test_db",
		DBUser:       "test_user",
		Runtime:      &model.Runtime{ID: 1, Name: "test runtime"},
		Pool:         &model.Pool{ID: 1, Name: "test pool"},
		Repository:   "https://example.repository.com",
		Branch:       "test_branch",
		Framework:    "test_framework",
		DocumentRoot: "/var/www/html",
		CreatedAt:    "2020-01-01",
		UpdatedAt:    "2020-01-01",
		Domains: []*model.Domain{
			{ID: 1, Name: "test-comain.com"},
		},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	response := NewApplicationDetailResponse(testUser, application)
	if response.Title != "Application Detail" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Application Detail" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Details) != 13 {
		t.Errorf("Details is not set properly. Got: %v", response.Details)
	}
}

// TestNewCreateApplicationResponse is a test function for the NewCreateApplicationResponse function.
// It tests the response generation.
func TestNewCreateApplicationResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	clients := &model.Clients{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	projects := &model.Projects{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	environments := &model.Environments{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	databases := &model.Databases{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	runtimes := &model.Runtimes{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	pools := &model.Pools{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	domains := &model.Domains{
		{ID: 1, Name: "test-domain.com", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test-domain2.com", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	response := NewCreateApplicationResponse(testUser, clients, projects, environments, databases, runtimes, pools, domains)
	if response.Title != "Create Application" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Create Application" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 13 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewUpdateApplicationResponse is a test function for the NewUpdateApplicationResponse function.
// It tests the response generation.
func TestNewUpdateApplicationResponse(t *testing.T) {
	application := &model.Application{
		ID:           1,
		Client:       &model.Client{ID: 1, Name: "test client"},
		Project:      &model.Project{ID: 1, Name: "test project"},
		Environment:  &model.Environment{ID: 1, Name: "test environment"},
		Database:     &model.Database{ID: 1, Name: "test database"},
		DBName:       "test_db",
		DBUser:       "test_user",
		Runtime:      &model.Runtime{ID: 1, Name: "test runtime"},
		Pool:         &model.Pool{ID: 1, Name: "test pool"},
		Repository:   "https://example.repository.com",
		Branch:       "test_branch",
		Framework:    "test_framework",
		DocumentRoot: "/var/www/html",
		CreatedAt:    "2020-01-01",
		UpdatedAt:    "2020-01-01",
		Domains: []*model.Domain{
			{ID: 1, Name: "test-comain.com"},
		},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	clients := &model.Clients{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	projects := &model.Projects{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	environments := &model.Environments{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	databases := &model.Databases{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	runtimes := &model.Runtimes{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	pools := &model.Pools{
		{ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test2", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	domains := &model.Domains{
		{ID: 1, Name: "test-domain.com", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01"},
		{ID: 2, Name: "test-domain2.com", CreatedAt: "2021-01-01", UpdatedAt: "2021-01-01"},
	}
	response := NewUpdateApplicationResponse(testUser, application, clients, projects, environments, databases, runtimes, pools, domains)
	if response.Title != "Update Application" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Update Application" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 13 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewApplicationListResponse is a test function for the NewApplicationListResponse function.
// It tests the response generation.
func TestNewApplicationListResponse(t *testing.T) {
	applications := &model.Applications{
		{
			ID:           1,
			Client:       &model.Client{ID: 1, Name: "test client"},
			Project:      &model.Project{ID: 1, Name: "test project"},
			Environment:  &model.Environment{ID: 1, Name: "test environment"},
			Database:     &model.Database{ID: 1, Name: "test database"},
			DBName:       "test_db",
			DBUser:       "test_user",
			Runtime:      &model.Runtime{ID: 1, Name: "test runtime"},
			Pool:         &model.Pool{ID: 1, Name: "test pool"},
			Repository:   "https://example.repository.com",
			Branch:       "test_branch",
			Framework:    "test_framework",
			DocumentRoot: "/var/www/html",
			CreatedAt:    "2020-01-01",
			UpdatedAt:    "2020-01-01",
			Domains: []*model.Domain{
				{ID: 1, Name: "test-comain.com"},
			},
		},
		{
			ID:           2,
			Client:       &model.Client{ID: 1, Name: "test client"},
			Project:      &model.Project{ID: 1, Name: "test project"},
			Environment:  &model.Environment{ID: 1, Name: "test environment"},
			Database:     &model.Database{ID: 1, Name: "test database"},
			DBName:       "test_db",
			DBUser:       "test_user",
			Runtime:      &model.Runtime{ID: 1, Name: "test runtime"},
			Pool:         &model.Pool{ID: 1, Name: "test pool"},
			Repository:   "https://example.repository.com",
			Branch:       "test_branch",
			Framework:    "test_framework",
			DocumentRoot: "/var/www/html/2",
			CreatedAt:    "2020-01-01",
			UpdatedAt:    "2020-01-01",
			Domains: []*model.Domain{
				{ID: 2, Name: "test-comain2.com"},
			},
		},
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view", "applications.create", "applications.update", "applications.delete", "clients.view", "projects.view", "environments.view", "databases.view", "runtimes.view", "pools.view", "domains.view"})
	response := NewApplicationListResponse(testUser, applications)
	if response.Title != "Application List" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Application List" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Header.Buttons) != 1 {
		t.Errorf("Header buttons are not set properly. Got: %v", response.Header.Buttons)
	}
}

// TestNewApplicationImportToEnvironmentFormResponse is a test function for the NewApplicationImportToEnvironmentFormResponse function.
// It tests the response generation.
func TestNewApplicationImportToEnvironmentFormResponse(t *testing.T) {
	environment := &model.Environment{
		ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01",
	}
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	response := NewApplicationImportToEnvironmentFormResponse(testUser, environment)
	if response.Title != "Import Application to Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Import Application to Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(response.Form.Items) != 2 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
}

// TestNewApplicationMappingToEnvironmentFormResponseWithoutHeaders is a test function for the NewApplicationMappingToEnvironmentFormResponse function.
// It tests the response generation when the headers are not set.
func TestNewApplicationMappingToEnvironmentFormResponseWithoutHeaders(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	environment := &model.Environment{
		ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01",
	}
	fileID := "test-identifier"
	csvData := [][]string{
		{"col1", "col2", "col3", "col4", "col5", "col6", "col7", "col8", "col9", "col10", "col11", "col12", "col13"},
	}
	headers := []string{}

	response := NewApplicationMappingToEnvironmentFormResponse(testUser, environment, fileID, csvData, headers)

	if response.Title != "Import Mapping to Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Import Mapping to Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	// 2 hidden input, 12 parameter for the header mapping, 12 parameter for custom value mapping
	if len(response.Form.Items) != 26 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
	if len(*response.Listing.Rows) != len(csvData) {
		t.Errorf("Invalid preview list length. Got: %d instead of %d", len(*response.Listing.Rows), len(csvData))
	}
}

// TestNewApplicationMappingToEnvironmentFormResponseWithHeaders is a test function for the NewApplicationMappingToEnvironmentFormResponse function.
// It tests the response generation when the headers are set.
func TestNewApplicationMappingToEnvironmentFormResponseWithHeaders(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	environment := &model.Environment{
		ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01",
	}
	fileID := "test-identifier"
	csvData := [][]string{
		{"col1", "col2", "col3", "col4", "col5", "col6", "col7", "col8", "col9", "col10", "col11", "col12", "col13"},
	}
	headers := []string{"header1", "project", "header3", "header4", "header5", "header6", "header7", "header8", "header9", "header10", "header11", "header12", "header13"}

	response := NewApplicationMappingToEnvironmentFormResponse(testUser, environment, fileID, csvData, headers)

	if response.Title != "Import Mapping to Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Import Mapping to Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	// 2 hidden input, 12 parameter for the header mapping, 12 parameter for custom value mapping
	if len(response.Form.Items) != 26 {
		t.Errorf("Form items are not set properly. Got: %v", response.Form.Items)
	}
	if len(*response.Listing.Rows) != len(csvData) {
		t.Errorf("Invalid preview list length. Got: %d instead of %d", len(*response.Listing.Rows), len(csvData))
	}
	// The project header is on index 1, it has to be the value for the project input. The project has to be under index 4 (env, file ids, project header id, project custom value).
	// Where the name of the form item is "project", then the Options where the value is "project" has to be selected. The others have to be not selected.
	for _, item := range response.Form.Items {
		if item.Options != nil {
			for _, option := range item.Options {
				if option.Value == "project" {
					// if the item name is "project" it has to be selected, otherwise not.
					if item.Name == "project" {
						if !option.Selected {
							t.Errorf("Invalid value for the project mapping. It supposed to be selected.")
						}
					} else {
						if option.Selected {
							t.Errorf("Invalid value for the project mapping. It supposed to be not selected.")
						}
					}
				} else {
					if option.Selected {
						t.Errorf("Invalid value for the project mapping. It supposed to be not selected.")
					}
				}
			}
		}
	}
}

// TestNewApplicationImportToEnvironmentListResponse is a test function for the NewApplicationImportToEnvironmentListResponse function.
// It tests the response generation.
func TestNewApplicationImportToEnvironmentListResponse(t *testing.T) {
	testUser := testhelper.GetUserWithAccessToResources(1, []string{"applications.view"})
	environment := &model.Environment{
		ID: 1, Name: "test", CreatedAt: "2020-01-01", UpdatedAt: "2020-01-01",
	}
	fileID := "test-identifier"
	importResults := parser.NewApplicationImportResult()
	// The first row contains a successful import.
	importResults[1] = &parser.ApplicationImportRow{
		ErrorMessage: "",
		RowData: map[string]string{
			"client":       "test client",
			"project":      "test project",
			"database":     "test database",
			"db_name":      "test_db",
			"db_user":      "test_user",
			"runtime":      "test runtime",
			"pool":         "test pool",
			"repository":   "https://example.repository.com",
			"branch":       "test_branch",
			"framework":    "test_framework",
			"documentRoot": "/var/www/html",
			"domains":      "test-domain.com",
		},
		Application: &model.Application{ID: 11},
	}
	// The second row contains an error message.
	importResults[2] = &parser.ApplicationImportRow{
		ErrorMessage: "test error message",
		RowData: map[string]string{
			"client":       "this is wrong",
			"project":      "test project",
			"database":     "test database",
			"db_name":      "test_db",
			"db_user":      "test_user",
			"runtime":      "test runtime",
			"pool":         "test pool",
			"repository":   "https://example.repository.com",
			"branch":       "test_branch",
			"framework":    "test_framework",
			"documentRoot": "/var/www/html",
			"domains":      "test-domain.com",
		},
		Application: nil,
	}

	response := NewApplicationImportToEnvironmentListResponse(testUser, environment, fileID, &importResults)
	if response.Title != "Import Application to Environment" {
		t.Errorf("Title is not set properly. Got: %s", response.Title)
	}
	if response.CurrentUser != testUser {
		t.Errorf("User is not set properly. Got: %v", response.CurrentUser)
	}
	if response.Header.Title != "Import Application to Environment" {
		t.Errorf("Header is not set properly. Got: %v", response.Header)
	}
	if len(*response.Listing.Rows) != 2 {
		t.Errorf("Listing length is not set properly. Got: %d", len(*response.Listing.Rows))
	}
}
